import { useParams } from "react-router-dom"
import { useContentMarkdown } from '../api/api'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import './Contents.css'
import ErrorPage from "./ErrorPage"

export const Content = () => {

    let { url } = useParams();
    let { data, isError } = useContentMarkdown(url)
    if (isError) return <ErrorPage code="NetorkError"></ErrorPage>

    return (
        <div>
            {/* <h2>{url}</h2> */}
            <div >
                <ReactMarkdown remarkPlugins={[remarkGfm]} children={data}></ReactMarkdown>
            </div>
        </div>

    )
}