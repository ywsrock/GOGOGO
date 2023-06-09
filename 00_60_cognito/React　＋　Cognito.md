# React　＋　Cognito

### 概要

既存のCognito利用して、認証機能を実装

> 1. Amplify　add auth は利用しない
> 2. Amplify　import auth は利用しない



## 設定方法

1. 下記のパッケージをインストール

   > ```json
   > @aws-amplify/ui-react": "^1.2.26",
   > "aws-amplify": "^5.0.20",
   > ```

2. index.jsファイルを下記のコード追加

   ```js
   import { Amplify } from 'aws-amplify';
   
   Amplify.configure({
     Auth: {
       // REQUIRED - Amazon Cognito Region
       region: 'apXXXXXXXXXXt-1',
       // OPTIONAL - Amazon Cognito User Pool ID
       userPoolId: 'XXXXXXXXXX',
       // OPTIONAL - Amazon Cognito Web Client ID (26-char alphanumeric string)
       userPoolWebClientId: 'XXXXX'
     },
   });
   ```

3. Amplify-uiの認証画面[AmplifyAuthenticator]を利用

   (追加はルートノートに追加)

   ```js
   import { AmplifyAuthenticator } from '@aws-amplify/ui-react';
   import { Auth } from 'aws-amplify';
   ・・・・・・・・・・・・・・
       return (
           <AmplifyAuthenticator>
               <Box sx={{ display: 'flex' }}>
                   {/* ログアウト */}
                   {/* <AmplifySignOut  variation="primary" style={{ marginRight: '0', marginLeft: 'auto' }} /> */}
                   <Button  size="small" variant="contained" color="secondary" style={{ marginRight: '0', marginLeft: 'auto' }}
                       onClick={() => { 
                           Auth.signOut().then(() => { 
                               window.location.href ="/";
                               // window.location.reload(false); 
                               }) }}>
                       signOut
                   </Button>
               </Box>
           </AmplifyAuthenticator>
   	)
   ```

4. Cognitoのユーザ情報利用

   ```js
   import { Auth } from 'aws-amplify';
   
   const [user, setUser] = useState();
       
       async function getUserInfo() {
           const AccountProfile = await Auth.currentSession()
           setUser(AccountProfile.idToken.payload);
           console.dir(AccountProfile.idToken.payload)
       }
       useEffect(() => {
           getUserInfo()
          
       }, [])
   ```

   