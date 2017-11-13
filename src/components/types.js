
export const TYPE_DOUBLE   = 1;
export const TYPE_FLOAT    = 2;
export const TYPE_INT64    = 3;
export const TYPE_UINT64   = 4;
export const TYPE_INT32    = 5;
export const TYPE_FIXED64  = 6;
export const TYPE_FIXED32  = 7;
export const TYPE_BOOL     = 8;
export const TYPE_STRING   = 9;
export const TYPE_GROUP     = 10;
export const TYPE_MESSAGE   = 11;
export const TYPE_BYTES   = 12;
export const TYPE_UINT32   = 13;
export const TYPE_ENUM   = 14;
export const TYPE_SFIXED32   = 15;
export const TYPE_SFIXED64   = 16;
export const TYPE_SINT32   = 17;
export const TYPE_SINT64   = 18;

export const INT_TYPES = {}

const ints = [
    TYPE_DOUBLE, TYPE_FLOAT, TYPE_INT64, TYPE_UINT64,
    TYPE_INT32, TYPE_FIXED64, TYPE_FIXED32, TYPE_UINT32,
    TYPE_SFIXED32, TYPE_SFIXED64, TYPE_SINT32, TYPE_SINT64
];
for (let i = 0; i <  ints.length; i++)  {
    INT_TYPES[ints[i]] = i;
}

export const getTypeName = (type_id) => {
    switch (type_id) {
        case TYPE_DOUBLE:  return "double";
        case TYPE_FLOAT:  return "float";
        case TYPE_INT64:  return "int64";
        case TYPE_UINT64:  return "uint64";
        case TYPE_INT32:  return "int32";
        case TYPE_FIXED64:  return "fixed64";
        case TYPE_FIXED32:  return "fixed32";
        case TYPE_BOOL:  return "bool";
        case TYPE_STRING:  return "string";
        case TYPE_GROUP: return "group";
        case TYPE_MESSAGE: return "message";
        case TYPE_BYTES: return "bytes";
        case TYPE_UINT32: return "uint32";
        case TYPE_ENUM: return "enum";
        case TYPE_SFIXED32: return "sfixed32";
        case TYPE_SFIXED64: return "sfixed64";
        case TYPE_SINT32: return "sint32";
        case TYPE_SINT64: return "sint64";
        default:
            return '???';
    }
};

export const getDefaultValue = (type_id, repeated, type_name, enums, types) => {
    if (repeated) {
        return [];
    }

    if (type_id in INT_TYPES) {
        return '0';
    }
    switch (type_id) {
        case TYPE_BOOL: //bool
            return 'false';
        case 11: //msg
            const type = types[type_name];
            return type.fields.map(f => getDefaultValue(f.type_id, f.is_repeated, f.type_name, enums, types));
        case 14:
            const e = enums[type_name].values;
            const keys = Object.keys(e);
            return  keys[0];
        default:
            return '';
    }
};