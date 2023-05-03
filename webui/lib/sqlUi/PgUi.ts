import { DataTypes, UITypes } from "~/utils/columnUtil";

export class PgUi {
    static getUiTypeFromDataType(dataType: DataTypes) {
        switch (dataType) {
            case DataTypes.Bigint:
            case DataTypes.Integer:
                return UITypes.Number;

            case DataTypes.Bool:
            case DataTypes.Boolean:
                return UITypes.Checkbox;

            case DataTypes.JsonB:
                return UITypes.JSON;

            case DataTypes.TimestampWithoutTimeZone:
            case DataTypes.Timestamp:
                return UITypes.DateTime;

            case DataTypes.Text:
            default:
                return UITypes.Text

        }
    }
    static getDataTypeFromUiType(type: UITypes) {
        switch (type) {

            case UITypes.Checkbox:
                return DataTypes.Boolean;

            case UITypes.SingleSelect:
            case UITypes.Password:
            case UITypes.Email:
                return DataTypes.Varchar;

            case UITypes.DateTime:
                return DataTypes.Timestamp;

            case UITypes.Number:
                return DataTypes.Bigint;

            case UITypes.JSON:
                return DataTypes.JsonB;


            case UITypes.Text:
            case UITypes.MultipleSelect:
            case UITypes.JSON:
            default:
                return DataTypes.Text;



        }
    }
}