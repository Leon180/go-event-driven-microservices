package enums

type EditType string

const (
	EditTypeNone   EditType = ""
	EditTypeCreate EditType = "create"
	EditTypeUpdate EditType = "update"
	EditTypeDelete EditType = "delete"
)

func (e EditType) ToString() string {
	return string(e)
}

func (e EditType) ToEditTypeCode() EditTypeCode {
	code, ok := EditTypeToEditTypeCodeMap[e]
	if !ok {
		return EditTypeCodeNone
	}
	return code
}

var EditTypeToEditTypeCodeMap = map[EditType]EditTypeCode{
	EditTypeNone:   EditTypeCodeNone,
	EditTypeCreate: EditTypeCodeCreate,
	EditTypeUpdate: EditTypeCodeUpdate,
	EditTypeDelete: EditTypeCodeDelete,
}

var EditTypes = []EditType{
	EditTypeNone,
	EditTypeCreate,
	EditTypeUpdate,
	EditTypeDelete,
}

type EditTypeCode int

const (
	EditTypeCodeNone EditTypeCode = iota
	EditTypeCodeCreate
	EditTypeCodeUpdate
	EditTypeCodeDelete
)

func (e EditTypeCode) ToEditType() EditType {
	if int(e) >= len(EditTypes) || int(e) < 0 {
		return EditTypeNone
	}
	return EditTypes[e]
}
