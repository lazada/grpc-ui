
const numberTypes ={
  double: true,
  float: true,
  int32: true,
  uint32: true,
  sint32: true,
  fixed32: true,
  sfixed32: true,
  int64: true,
  uint64: true,
  sint64: true,
  fixed64: true,
  sfixed64: true,
};

export function isNumberType(type: string): boolean {
  if (numberTypes[type]) {
    return true;
  }

  return false;
}
