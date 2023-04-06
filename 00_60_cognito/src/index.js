import { Amplify } from 'aws-amplify';
import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  RouterProvider
} from "react-router-dom";
import './index.css';
import reportWebVitals from './reportWebVitals';
import { router } from './router/router';

Amplify.configure({
  Auth: {
    // REQUIRED - Amazon Cognito Region
    region: 'ap-northeast-1',
    // OPTIONAL - Amazon Cognito User Pool ID
    // userPoolId: 'ap-northeast-1_rfx4sP2lH',
    userPoolId: 'ap-northeast-1_r1SlHjWSb',
    // OPTIONAL - Amazon Cognito Web Client ID (26-char alphanumeric string)
    // userPoolWebClientId: '3iu7lb1u54pi9q5ptq4v9kpfdu'
    userPoolWebClientId: '6654cm8banrk0agfrf6jmc6ejb'
  },
  // API:{
  //   endpoints: [
  //     {
  //       name:"DemoApi",
  //       endpoint:"https://lnke3aknkl.execute-api.ap-northeast-1.amazonaws.com/AfDemoApiState"
  //     }
  //   ]
  // }
});



const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
