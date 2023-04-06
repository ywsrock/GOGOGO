import React from 'react';
import { createBrowserRouter } from "react-router-dom";
import CustomerRequestInfo from '../component/customerRequestInfo';
import HistoryDetailsList from '../component/historyDetailsList';
import Layout from '../component/layout';
import NotificationSetting from '../component/notification-setting';
import ReceptionStatusList from '../component/receptionStatusList';
import ShowCustomerRequestInfo, { loader } from '../component/showcustomerRequestInfo';
import SystemSetting from '../component/system-setting';
import { UserInfo } from '../component/userinfo';
// import AmplifyAuthenticator from '@aws-amplify/ui-react';

export const router = createBrowserRouter([
  // {
  //   path: "/",
  //   element: <SignIn></SignIn>,
  // },
  {
    path: "/",
    element: <Layout></Layout>,
    children:[
      {
      index:true,
      element:<ReceptionStatusList></ReceptionStatusList>
    },{
      path:"ShowCustomerRequestInfo/:id",
      loader:loader,
      element:<ShowCustomerRequestInfo></ShowCustomerRequestInfo>
    },
    {
      path:"CustomerRequestInfo",
      element:<CustomerRequestInfo></CustomerRequestInfo>
    },
    {
      path:"userInfo",
      element:<UserInfo></UserInfo>
    },
    {
      path:"historyDetails",
      element:<HistoryDetailsList></HistoryDetailsList>
    },
    {
      path:"receptionStatusList",
      element:<ReceptionStatusList></ReceptionStatusList>
    },
    {
      path:"notificationSetting",
      element:<NotificationSetting></NotificationSetting>
    },
    {
      path:"systemSetting",
      element:<SystemSetting></SystemSetting>
    },
    ]
  },

]);
