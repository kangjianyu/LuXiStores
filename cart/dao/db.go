package cart_dao

import "LuXiStores/common"

var DB iBD = dbimpl{}

type iBD interface {
	GetGoodsCartList(uid uint64,count uint64,start uint64) ([]GoodsCartInfo,error)
	DelGoodsCart(uid uint64,productId uint64) error
	UpdateGoodsCart(uid uint64,id uint64,quantity uint64) error
	AddGoodsCart(uid uint64,productId uint64,quantity uint64) error

}


type dbimpl struct {

}

func (dbimpl) GetGoodsCartList(uid uint64, count uint64,start uint64) ([]GoodsCartInfo, error) {
	tablename := (&GoodsCartInfo{}).TableName()
	info := []GoodsCartInfo{}
	err := common.MysqlClient.GetGoodsCartList(tablename,uid,count,start,&info).Error
	return info,err
}

func (dbimpl) DelGoodsCart(uid uint64, productId uint64) error {
	tablename := (&GoodsCartInfo{}).TableName()
	err := common.MysqlClient.DelGoodsCart(tablename,uid,productId).Error
	return err
}

func (dbimpl) UpdateGoodsCart(uid uint64, id uint64, quantity uint64) error {
	tablename := (&GoodsCartInfo{}).TableName()
	err := common.MysqlClient.UpdateGoodsCartCount(tablename,id,quantity,uid).Error
	return err
}

func (dbimpl) AddGoodsCart(uid uint64, productId uint64, quantity uint64) error {
	tablename := (&GoodsCartInfo{}).TableName()
	err := common.MysqlClient.AddGoodsCart(tablename,uid,quantity,productId).Error
	return err

}

