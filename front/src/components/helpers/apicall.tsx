import axios from "axios";
import {Create, Insert, Select} from "../Quering/index";

const BASE_URL = "http://localhost:4200/api"

export const createTable = async (data: Create): Promise<any> => {
    data.attribs.set("TABLE", data.table);
    const json = Object.fromEntries(data.attribs)
    const url = `${BASE_URL}/create-table`
    console.log(json);
    return await axios.post(url,json);
}


export const insertIntoTable = async (data: Insert): Promise<any> => {
    const table = data.table;
    const url = `${BASE_URL}/insert-record/${table.toLowerCase()}`
    const json = Object.fromEntries(data.values)
    console.log(json)
    return await axios.post(url,json);
}

export const selectToTable = async (data: Select): Promise<any> => {
    const table = data.table;
    const url = `${BASE_URL}/records-filtered/${table.toLowerCase()}`
    const json = data.attribs
    console.log(json)
    return await axios.post(url, {
        "data": json
    });
}

export const allTables = async (): Promise<any> => {
    const url =`${BASE_URL}/tables` 
    return await axios.get(url);
}