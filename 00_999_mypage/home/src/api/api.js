import useSWR from 'swr'
import axios from 'axios'

const API = axios.create({
    baseURL: 'http://localhost:8003',
    timeout: 1000,
    responseType: "json",
});


export const fetcher2 = (url) => {
    let u = require("../markdown/" + url);
    return fetch(u).then(res => res.text())
};

const fetcher = async (url) => {
    try {
        const response = await API.get(url)
        return response.data;
    } catch (error) {
        return error;
    }
}


const swrOption = {
    suspense: true,
    // revalidateOnMount: true,
    revalidateOnFocus: true,
    revalidateOnReconnect: true,
    shouldRetryOnError: true,
    errorRetryCount: 3,
}

// Home infomation
export const useHome = (url) => {
    const { data, error } = useSWR(url, fetcher, swrOption)
    return {
        data: data,
        isLoading: !error && !data,
        isError: error || data.code
    }
}

// About infomation
export const useAbout = (url) => {
    const { data, error } = useSWR(url, fetcher, swrOption)
    return {
        data: data,
        isLoading: !error && !data,
        isError: error || data.code
    }
}

export const useContentsList = (url) => {
    const { data, error } = useSWR(url, fetcher, swrOption)
    return {
        data: data,
        isLoading: !error && !data,
        isError: error || data.code
    }
}

// load markdown file
export const useContentMarkdown = (url) => {
    const { data, error } = useSWR(url, fetcher2, swrOption)
    return {
        data: data,
        isLoading: !error && !data,
        isError: error || data.code
    }
}