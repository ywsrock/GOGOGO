import useSWR from 'swr';
import useSWRImmutable from 'swr/immutable';
import axios from 'axios'
import { isUndefined, trim } from 'lodash';
import { WordDataInit, HistoryDataInit, StsDataInit } from '../state/State'

export const API = axios.create({
    baseURL: 'http://localhost:8003',
    timeout: 1000,
    responseType: "json",
});

//WORD
const DicBase = "/tool/dic/v1/c"

export const fetcher2 = (url) => {
    let u = require("../markdown/" + url);
    return fetch(u).then(res => res.text())
};

export const AxiosFetcher = async (url) => {
    try {
        // console.log(url)
        const response = await API.get(url)
        return response.data;
    } catch (error) {
        return error;
    }
}

// SWROption info
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
    const { data, error } = useSWR(url, AxiosFetcher, swrOption)
    return {
        data: data,
        isLoading: !error && !data,
        isError: error || data.code
    }
}

// About infomation
export const useAbout = (url) => {
    const { data, error } = useSWR(url, AxiosFetcher, swrOption)
    return {
        data: data,
        isLoading: !error && !data,
        isError: error || data.code
    }
}

export const useContentsList = (url) => {
    const { data, error } = useSWR(url, AxiosFetcher, swrOption)
    return {
        data: data,
        isLoading: !error && !data,
        isError: error || data.code
    }
}

// Load markdown file
export const useContentMarkdown = (url) => {
    const { data, error } = useSWR(url, fetcher2, swrOption)
    return {
        data: data,
        isLoading: !error && !data,
        isError: error || data.code
    }
}

// Load dictional infomation
export const useDictionalAll = (url1) => {
    let url = `${DicBase}/sts`
    const { data, error } = useSWR(url, AxiosFetcher, swrOption)
    return {
        StsData: data,
        isLoading: !error & !data,
        isError: error || data.code
    }
}

//Load dictional word
export const useDictionalHistory = () => {
    let url = `${DicBase}/historyAll`
    let { data, error } = useSWR(url, AxiosFetcher, swrOption)
    if (isUndefined(data)) {
        data = {
            code: "",
            ...HistoryDataInit,
        }
    }
    // data = isUndefined(data) ? HistoryDataInit : data
    return {
        HistoryData: data,
        HistoryIsLoading: !error & !data,
        HistoryIsError: error || data.code
    }
}

//Load dictional word
export const useDictionalWord = (key) => {
    let url = trim(key) === "" ? DicBase : `${DicBase}/?key=${key}`
    console.log("Url--->" + url)
    let { data, error } = useSWRImmutable(key !== "" ? url : null, AxiosFetcher, { suspense: true })
    // let { data, mutate, error } = useSWR(url, AxiosFetcher, swrOption)
    if (isUndefined(data)) {
        data = {
            code: "",
            ...WordDataInit,
        }
    }
    // data = isUndefined(data) ? WordDataInit : data
    return {
        WordData: data,
        WordIsLoading: !error & !data,
        WordIsError: error || data.code,
    }
}

export const useDictionalSts = () => {
    let url = `${DicBase}/sts`
    let { data, error } = useSWR(url, AxiosFetcher, swrOption)
    if (isUndefined(data)) {
        data = {
            code: "",
            ...StsDataInit,
        }
    }
    // data = isUndefined(data) ? StsDataInit : data
    return {
        StsData: data,
        StsIsLoading: !error & !data,
        StsIsError: error || data.code
    }
}

export const QueryWord = (url) => useDictionalAll(url)