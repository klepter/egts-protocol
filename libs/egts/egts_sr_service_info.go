package egts

import (
	"bytes"
	"fmt"
)

//SrServiceInfo структура подзаписи типа EGTS_SR_SERVICE_INFO, которая применяется телематической
//платформой для информирования АС о состоянии какого-либо сервиса.
type SrServiceInfo struct {
	ServiceType       		uint8  `json:"ST"`
	ServiceStatement  		uint8  `json:"SST"`
	ServiceAttribute        string `json:"SRVA"`
	ServiceRoutingPriority  string `json:"SRVRP"`
}

//Decode разбирает байты в структуру подзаписи
func (s *SrServiceInfo) Decode(content []byte) error {
	var (
		err error
		flags byte
	)
	buf := bytes.NewBuffer(content)

	if s.ServiceType, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить тип сервиса: %v", err)
	}

	if s.ServiceStatement, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить состояние сервиса: %v", err)
	}

	if flags, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Не удалось получить параметры сервиса: %v", err)
	}

	flagBits := fmt.Sprintf("%08b", flags)
	e.ServiceAttribute = flagBits[:1]
	e.ServiceRoutingPriority = flagBits[6:]

	return err
}

//Encode преобразовывает подзапись в набор байт
func (s *SrServiceInfo) Encode() ([]byte, error) {
	var (
		result []byte
		flags  uint64
		err    error
	)
	buf := new(bytes.Buffer)

	if err = buf.WriteByte(s.ServiceType); err != nil {
		return result, fmt.Errorf("Не удалось записать тип сервиса: %v", err)
	}

	if err = buf.WriteByte(s.ServiceStatement); err != nil {
		return result, fmt.Errorf("Не удалось записать состояние сервиса: %v", err)
	}

	flags, err = strconv.ParseUint(e.ServiceAttribute, e.ServiceRoutingPriority, 2, 8)
	if err = buf.WriteByte(uint8(flags)); err != nil {
		return result, fmt.Errorf("Не удалось записать байт параметров сервиса: %v", err)
	}

	result = buf.Bytes()
	return result, err
}

//Length получает длинну закодированной подзаписи
func (s *SrServiceInfo) Length() uint16 {
	var result uint16

	if recBytes, err := s.Encode(); err != nil {
		result = uint16(0)
	} else {
		result = uint16(len(recBytes))
	}

	return result
}
