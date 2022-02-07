'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

const graphql = require('graphql');
const visitorPluginCommon = require('@graphql-codegen/visitor-plugin-common');
const javaCommon = require('@graphql-codegen/java-common');
const path = require('path');

class JavaResolversVisitor extends visitorPluginCommon.BaseVisitor {
    constructor(rawConfig, _schema, defaultPackageName) {
        super(rawConfig, {
            mappers: visitorPluginCommon.transformMappers(rawConfig.mappers || {}),
            package: rawConfig.package || defaultPackageName,
            defaultMapper: visitorPluginCommon.parseMapper(rawConfig.defaultMapper || 'Object'),
            className: rawConfig.className || 'Resolvers',
            listType: rawConfig.listType || 'Iterable',
            scalars: visitorPluginCommon.buildScalars(_schema, rawConfig.scalars, javaCommon.JAVA_SCALARS, 'Object'),
        });
        this._includeTypeResolverImport = false;
    }
    getImports() {
        const mappersImports = this.mappersImports();
        const allImports = [...mappersImports];
        if (this._includeTypeResolverImport) {
            allImports.push('graphql.schema.TypeResolver');
        }
        allImports.push('graphql.schema.DataFetcher');
        return allImports.map(i => `import ${i};`).join('\n') + '\n';
    }
    mappersImports() {
        return Object.keys(this.config.mappers)
            .map(typeName => this.config.mappers[typeName])
            .filter((m) => m.isExternal)
            .map(m => m.source);
    }
    getTypeToUse(type) {
        if (this.scalars[type.name.value]) {
            return this.scalars[type.name.value];
        }
        else if (this.config.mappers[type.name.value]) {
            return this.config.mappers[type.name.value].type;
        }
        return this.config.defaultMapper.type;
    }
    getPackageName() {
        return `package ${this.config.package};\n`;
    }
    wrapWithClass(content) {
        return new javaCommon.JavaDeclarationBlock()
            .access('public')
            .asKind('class')
            .withName(this.config.className)
            .withBlock(visitorPluginCommon.indentMultiline(content)).string;
    }
    UnionTypeDefinition(node) {
        this._includeTypeResolverImport = true;
        return new javaCommon.JavaDeclarationBlock()
            .access('public')
            .asKind('interface')
            .withName(this.convertName(node.name))
            .extends(['TypeResolver'])
            .withComment(node.description).string;
    }
    InterfaceTypeDefinition(node) {
        this._includeTypeResolverImport = true;
        return new javaCommon.JavaDeclarationBlock()
            .access('public')
            .asKind('interface')
            .withName(this.convertName(node.name))
            .extends(['TypeResolver'])
            .withComment(node.description)
            .withBlock(node.fields.map(f => visitorPluginCommon.indent(f(true))).join('\n')).string;
    }
    ObjectTypeDefinition(node) {
        return new javaCommon.JavaDeclarationBlock()
            .access('public')
            .asKind('interface')
            .withName(this.convertName(node.name))
            .withComment(node.description)
            .withBlock(node.fields.map(f => visitorPluginCommon.indent(f(false))).join('\n')).string;
    }
    FieldDefinition(node, key, _parent) {
        return (isInterface) => {
            const baseType = visitorPluginCommon.getBaseTypeNode(node.type);
            const typeToUse = this.getTypeToUse(baseType);
            const wrappedType = javaCommon.wrapTypeWithModifiers(typeToUse, node.type, this.config.listType);
            if (isInterface) {
                return `default public DataFetcher<${wrappedType}> ${node.name.value}() { return null; }`;
            }
            else {
                return `public DataFetcher<${wrappedType}> ${node.name.value}();`;
            }
        };
    }
}

const plugin = async (schema, documents, config, { outputFile }) => {
    const relevantPath = path.dirname(path.normalize(outputFile));
    const defaultPackageName = javaCommon.buildPackageNameFromPath(relevantPath);
    const visitor = new JavaResolversVisitor(config, schema, defaultPackageName);
    const printedSchema = graphql.printSchema(schema);
    const astNode = graphql.parse(printedSchema);
    const visitorResult = graphql.visit(astNode, { leave: visitor });
    const mappersImports = visitor.getImports();
    const packageName = visitor.getPackageName();
    const blockContent = visitorResult.definitions.filter(d => typeof d === 'string').join('\n');
    const wrappedContent = visitor.wrapWithClass(blockContent);
    return [packageName, mappersImports, wrappedContent].join('\n');
};

exports.plugin = plugin;
//# sourceMappingURL=index.cjs.js.map
