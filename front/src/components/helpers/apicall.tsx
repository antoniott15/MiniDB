import axios from "axios";
import {Create, Insert, Select} from "../Quering/index";

const BASE_URL = "http://localhost:4200/api"

export const createTable = async (data: Create): Promise<any> => {
    data.attribs.set("TABLE", data.table);
    const json = Object.fromEntries(data.attribs)
    const url = `${BASE_URL}/create-table`
    return await axios.post(url,json);
}


export const insertIntoTable = async (data: Insert): Promise<any> => {
    const table = data.table;
    const url = `${BASE_URL}/insert-record/${table}`
    const json = Object.fromEntries(data.values)
    return await axios.post(url,json);
}

export const selectToTable = async (data: Select): Promise<any> => {
    const table = data.table;
    const url = `${BASE_URL}/records-filtered/${table}`
    const json = data.attribs
    return await axios.post(url, {
        "data": json
    });
}

export const allTables = async (): Promise<any> => {
    const url =`${BASE_URL}/tables` 
    return await axios.get(url);
}