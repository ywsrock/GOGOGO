import {
    createBrowserRouter,
} from "react-router-dom";
import App from "../App";
import ErrorPage from '../component/ErrorPage'
import { Home } from "../component/Home";
import { About } from "../component/About";
import { News } from "../component/News";
import { Tools } from "../component/Tools";
import { Content } from '../component/Contents'

const router = createBrowserRouter(
    [
        {
            path: "/",
            element: <App />,
            errorElement: <ErrorPage message="404" />,
            children: [
                {
                    index: true,
                    element: <Home />
                },
                {
                    index: true,
                    path: "/home",
                    element: <Home />
                },
                {
                    path: "/about",
                    element: <About />
                },
                {
                    path: "/news",
                    element: <News />
                },
                {
                    path: "/tool",
                    element: <Tools />
                },
            ]
        },
        {
            path: "/Contents/:url",
            element: <Content />
        }
    ])

export default router;