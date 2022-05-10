package warehouse

const (
	// --------------------Константы кодов склада для идентификации--------------------
	GeneralPickZoneCode = 7000
	GeneralReserveCode  = 7010

	FoilPickZoneCode = 7100
	FoilReserveCode  = 7110

	NetsPickZoneCode = 7200
	NetsReserveCode  = 7210

	// --------------------Константы кодов особой ситуации--------------------
	NoProductInPlace        = "BIDU" // В случае отсутствия продукта на месте
	NotEnoughProductInPlace = "BIDF" // В случае, когда невозможно взять требуемое системой кол-во продукта. Например 1 шт. при минимальной единице измерения 1 блк.
)
