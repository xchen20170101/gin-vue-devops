package datas

import (
	"os"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Carbines = []gormadapter.CasbinRule{
	{PType: "p", V0: "888", V1: "/base/login", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/register", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getApiById", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/deleteApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/updateApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getAllApis", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/createAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/deleteAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/getAuthorityList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/setDataAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/updateAuthority", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/authority/copyAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{PType: "p", V0: "888", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{PType: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
	{PType: "p", V0: "888", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{PType: "p", V0: "888", V1: "/casbin/casbinTest/:pathParam", V2: "GET"},
	{PType: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{PType: "p", V0: "888", V1: "/system/getSystemConfig", V2: "POST"},
	{PType: "p", V0: "888", V1: "/system/setSystemConfig", V2: "POST"},
	{PType: "p", V0: "888", V1: "/system/getServerInfo", V2: "POST"},
	{PType: "p", V0: "888", V1: "/autoCode/createTemp", V2: "POST"},
	{PType: "p", V0: "888", V1: "/autoCode/getTables", V2: "GET"},
	{PType: "p", V0: "888", V1: "/autoCode/getDB", V2: "GET"},
	{PType: "p", V0: "888", V1: "/autoCode/getColumn", V2: "GET"},
	{PType: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/email/emailTest", V2: "POST"},
	{PType: "p", V0: "888", V1: "/k8sCluster/createK8sCluster", V2: "POST"},
	{PType: "p", V0: "888", V1: "/k8sCluster/findK8sCluster", V2: "GET"},
	{PType: "p", V0: "888", V1: "", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sNamespaces/findK8sNamespaces", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sNamespaces/getK8sNamespacesList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sDeployments/getK8sDeploymentList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/k8sPods/getK8sPodsList", V2: "GET"},
	{PType: "p", V0: "8881", V1: "/base/login", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/user/register", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/createApi", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/getApiList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/getApiById", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/deleteApi", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/updateApi", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/getAllApis", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/authority/createAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/authority/deleteAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/authority/getAuthorityList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/authority/setDataAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getMenu", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getMenuList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/addBaseMenu", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/addMenuAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getMenuAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/updateBaseMenu", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getBaseMenuById", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/user/changePassword", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/user/getUserList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/user/setUserAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/casbin/updateCasbin", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/system/getSystemConfig", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/system/setSystemConfig", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/base/login", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/user/register", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/createApi", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/getApiList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/getApiById", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/deleteApi", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/updateApi", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/getAllApis", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/authority/createAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/authority/deleteAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/authority/getAuthorityList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/authority/setDataAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getMenu", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getMenuList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/addBaseMenu", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/addMenuAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getMenuAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/updateBaseMenu", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getBaseMenuById", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/user/changePassword", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/user/getUserList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/user/setUserAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/casbin/updateCasbin", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/system/getSystemConfig", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/system/setSystemConfig", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/autoCode/createTemp", V2: "POST"},
}

func InitCasbinModel(db *gorm.DB) {
	if err := db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("p_type = ? AND v0 IN ?", "p", []string{"888", "8881", "9528"}).Find(&[]gormadapter.CasbinRule{}).RowsAffected == 142 {
			color.Danger.Println("casbin_rule表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&Carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	}); err != nil {
		color.Warn.Printf("[Mysql]--> casbin_rule 表的初始数据失败,err: %v\n", err)
		os.Exit(0)
	}
}
