export enum UITypes {
    ID = 'id',
    Text = 'Text',
    Email = 'Email',
    SingleSelect = 'SingleSelect',
    MultipleSelect = 'MultipleSelect',
    Checkbox = 'Checkbox',
    DateTime = 'DateTime',
    Password = 'Password',
    Number = 'Number',
    JSON = 'JSON',
    Other = 'Other',
}

export enum DataTypes {
  Text = 'text',
  Varchar = 'varchar',
  Bool = 'bool',
  Boolean = 'boolean',
  Timestamp = 'timestamp',
  TimestampWithoutTimeZone = 'timestamp without time zone',
  Bigint = 'bigint',
  Integer = 'integer',
  JsonB = 'jsonb'
}

export const uiTypes = [
  {
    name: UITypes.Text,
    icon: iconMap.text,
  },
  {
    name: UITypes.Number,
    icon: iconMap.number,
  },
  {
    name: UITypes.Email,
    icon: iconMap.email,
  },
  {
    name: UITypes.SingleSelect,
    icon: iconMap.singleSelect,
  },
  {
    name: UITypes.MultipleSelect,
    icon: iconMap.multiSelect,
  },
  {
    name: UITypes.Checkbox,
    icon: iconMap.check,
  },
  {
    name: UITypes.DateTime,
    icon: iconMap.datetime,
  },
  {
    name: UITypes.Password,
    icon: iconMap.passwordChange,
  },
  {
    name: UITypes.JSON,
    icon: iconMap.json,
  },
  {
    name: UITypes.Other,
    icon: iconMap.other,
  },
];
