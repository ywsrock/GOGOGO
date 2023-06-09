export const RowsPerPage = 10

export const MainMenu = [
    {
        name:"依頼一覧",
        component:"userList",
        rout:"/receptionStatusList"
    },
    {
        name:"対応履歴",
        component:"History",
        rout:"/historyDetails"
    },
    {
        name:"応対登録",
        component:"Main",
        rout:"/CustomerRequestInfo"
        
    },
    {
        name:"応対管理",
        component:"Main",
        rout:"/"
    },
]


export const SubMenu = [
    {
        name:"ログイン情報",
        component:"UserInfo",
        rout:"/userInfo"
    },
    {
        name:"通知設定",
        component:"NotificationSetting",
        rout:"/notificationSetting"
    },
    {
        name:"システム設定",
        component:"SystemSetting",
        rout:"/systemSetting"
    },
    {
        name:"設定",
        component:"",
        rout:"/"
    },
]