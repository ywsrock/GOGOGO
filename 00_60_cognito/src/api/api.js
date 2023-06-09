import axios from "axios";

const API = axios.create({
	baseURL: process.env.REACT_APP_SERVER_HOST_IP,
	timeout: 5000,
});

//新規作成
export const NewRequest = async (PostData) => {
	const { data } = await API.post("/", PostData)
	console.log(data)
	return data
}

//指定IDで更新
export const Updatehis = async (PostData) => {
	const { data } = await API.put("/updatehis", PostData)
	console.log(data)
	return data
}

//全件検索
export const ShowHistory = async () => {
	const { data } = await API.get("/historys")
	return data
}

//1件検索
export const GetHistory = async (id) => {
	const { data } = await API.get("/history", {
		params: {
			"recordID": id
		}
	})
	return data
}