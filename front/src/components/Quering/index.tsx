import React, { useState } from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import { Typography } from '@material-ui/core';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import SearchIcon from '@material-ui/icons/Search';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { createTable, insertIntoTable, selectToTable } from "../helpers/apicall";

var counter = 0;
export const CounterContext = React.createContext(counter);

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            flexGrow: 1,
            margin: "2em"
        },
        paper: {
            padding: theme.spacing(2),
            textAlign: 'center',
            color: theme.palette.text.secondary,
        },
        button: {
            margin: theme.spacing(1),
            color: theme.palette.secondary.main,
            decomposeColor: theme.palette.primary.main,
        },
        gridButton: {
            display: "flex",
            justifyContent: "flex-end"
        },
        table: {
            minWidth: 500,
        },
        tableW:{
            maxHeight: 440,
        }
    }),
);

export interface Insert {
    query: string,
    table: string,
    key: number,
    values: Map<string, string>,
}

export interface Select {
    query: string,
    table: string,
    attribs: string[],
}

export interface Create {
    query: string,
    table: string,
    key: string | undefined;
    attribs: Map<string, string>,
}


const deleteQuotes = (q: string): string => {
    return q.split(",")[0];
}

const removeParenthesis = (q: string): string[] => {
    let newWord: string = "";
    for (let elements of q) {
        if (elements !== "(" && elements !== ")") {
            newWord += elements;
        }
    };
    const result = newWord.split(",");
    return result;
}



const Quering = () => {
    const classes = useStyles();
    const [value, setValue] = useState<string>("")
    const [notFound, setNotFound] = useState<boolean>(false);
    const [notFoundFROM, setNotFoundFROM] = useState<boolean>(false);
    const [valuesNotEqual, setValuesNotEqual] = useState<boolean>(false);
    const [notKey, setNotKey] = useState<boolean>(false);
    const [createdTable, setCreatedTable] = useState<boolean>(false);
    const [insertingRecord, setInsertingRecord] = useState<boolean>(false);
    const [tableExist, setTableExist] = useState<boolean>(false);
    const [errorInsert, errorInserting] = useState<boolean>(false);
    const [table, selectedTable] = useState<boolean>(false);
    const [errSelct, errInSelecting] = useState<boolean>(false);
    const [headers, setHeaders] = useState<string[]>([]);
    const [body, setBody] = useState<any[]>([]);

   

    const desactiveBooleans = (t: boolean) => {
        setNotFound(t)
        setNotFoundFROM(t)
        setValuesNotEqual(t);
        setNotKey(t);
        setInsertingRecord(t);
        setCreatedTable(t);
        setTableExist(t);
        errorInserting(t);
        selectedTable(t);
        errInSelecting(t);
    }


    const callAxios = (): void => {
        const query: string[] = value.split(" ")
        if (query[0] === "insert") {
            desactiveBooleans(false)
            const insertQuery: Insert | undefined = insert(query.slice(1));
            if (insertQuery !== undefined) {
                insertIntoTable(insertQuery).then((_) => { setInsertingRecord(true) }).catch(e => {
                    desactiveBooleans(false);
                     errorInserting(true)});
            }
        } else if (query[0] === "create") {
            desactiveBooleans(false)
            const createQuery: Create | undefined = create(query.slice(1));
            if (createQuery !== undefined) {
                createTable(createQuery).then((_) => { 
                    setCreatedTable(true) 
                    counter++;
                }).catch(e => {
                    console.log(e)
                    desactiveBooleans(false);
                    setTableExist(true)});
            }
        } else if (query[0] === "select") {
            desactiveBooleans(false)
            const selectQuery: Select | undefined = select(query.slice(1))
            if (selectQuery !== undefined) {
                selectToTable(selectQuery).then((res) => {
                    selectedTable(true);
                    const data = res.data.data;
                    setHeaders(data[0].Headers);
                    var arr = new Array<any>();
                    for (const elements of data) {
                        arr.push(elements.Attribs)
                    }
                    setBody(arr);
                }).catch(e => {
                    desactiveBooleans(false);
                    errInSelecting(true)
                })
            }
        } else {
            setNotFound(true)
            return;
        }
    }

    const create = (query: string[]): Create | undefined => {
        const table = query[1];
        var word: string = ""
        for (let elements of query) {
            word += elements;
        }
        const attribs = new Map<string, string>();
        word.split(",").forEach((values, i) => {
            if (i === 0) {
                const at = values.split("(")[1].split(":");
                attribs.set(at[0].replace("\n", "").toUpperCase(), at[1])
            } else {
                const at = values.split(":");
                attribs.set(at[0].replace("\n", "").toUpperCase(), at[1]);
            }
        })
        attribs.delete(");");

        if (attribs.get("KEY") === undefined) {
            setNotKey(true)
            return undefined;
        }

        var result: Create = {
            query: "create",
            table: table,
            key: attribs.get("KEY"),
            attribs: attribs
        }
        return result;
    }

    const insert = (query: string[]): Insert | undefined => {
        const into = query.findIndex((values) => values === "into");
        const values = query.findIndex((values) => values === "values");
        const table = query.slice(into + 1, values)[0];
        const records = removeParenthesis(query.slice(query.findIndex((values) => values === table) + 1, values)[0]);

        const value = removeParenthesis(query.slice(values)[1].split(";")[0]);
        if (value.length !== records.length) {
            setValuesNotEqual(true);
            return;
        }

        let Result = new Map<string, string>();
        records.forEach((v, i) => {
            Result.set(v.toUpperCase(), value[i]);
        })

        var result: Insert = {
            query: "insert",
            table: table,
            key: Number(Result.get("KEY")),
            values: Result,
        }

        return result;
    }

    const select = (query: string[]): Select | undefined => {
        var newQuery: Select = {
            query: "select",
            table: "",
            attribs: new Array<string>(),
        }
        const fromIndex = query.findIndex((values) => values === "from");
        if (fromIndex === -1) {
            setNotFoundFROM(true)
            return undefined;
        }
        const table = query[fromIndex + 1];
        const newTable = table.split(";");
        newQuery.table = newTable[0];

        if (query[0] === "*") {
            newQuery.attribs.push("*");
        } else {
            const values = new Array<string>();
            for (let elements of query.slice(0, fromIndex)) {
                values.push(deleteQuotes(elements).toUpperCase())
            }
            newQuery.attribs.push.apply(newQuery.attribs, values);
        }

        return newQuery;
    }



    return (
        <div className={classes.root}>
            <Grid container spacing={3}>
                <Grid item xs={3}>
                    <Paper className={classes.paper}>
                        <Typography variant="subtitle1">
                            En la siguiente sección puedes escribir tus querys tales como:
                            <ul>
                                <li>Insert</li>
                                <li>Select</li>
                                <li>Create</li>
                            </ul>
                        </Typography>
                    </Paper>
                </Grid>
                <Grid item xs={9} sm={9}>
                    <Paper className={classes.paper}>
                        <TextField id="standard-basic" label="Create, select or insert" placeholder="Querys goes here" fullWidth rows={5} multiline value={value.toLowerCase()} onChange={(e) => setValue(e.target.value)} />
                        {notFound ? <Typography variant="subtitle2" color="error"> Query rechazado, el query debería comenzar con INSERT, CREATE O SELECT </Typography> : ""}
                        {notFoundFROM ? <Typography variant="subtitle2" color="error"> FROM no encontrado en select query </Typography> : ""}
                        {valuesNotEqual ? <Typography variant="subtitle2" color="error"> Keys y Values no hacen match en el tamaño </Typography> : ""}
                        {notKey ? <Typography variant="subtitle2" color="error"> La tabla creada necesita una KEY para poder crearse </Typography> : ""}
                        <Grid className={classes.gridButton}>

                            <Button
                                variant="contained"
                                color="primary"
                                onClick={callAxios}
                                className={classes.button}
                                endIcon={<SearchIcon />}
                            >
                                Query
                        </Button>
                        </Grid>
                    </Paper>
                </Grid>
                <Grid item xs={9} sm={12}>
                    <Paper className={classes.paper}>
                        {createdTable ? <Typography variant="subtitle2" color="inherit"> Se ha creado la tabla satisfactoriamente </Typography> : ""}
                        {tableExist ? <Typography variant="subtitle2" color="error"> Tabla ya existente o no se reconoce </Typography> : ""}
                        {insertingRecord ? <Typography variant="subtitle2" color="inherit"> Se ha insertado un nuevo record en la tabla especificada correctamente </Typography> : ""}
                        {errorInsert ? <Typography variant="subtitle2" color="error"> Error al intentar insertar un nuevo record </Typography> : ""}
                        {errSelct ? <Typography variant="subtitle2" color="error"> Error al seleccionar una tabla </Typography> : ""}
                        {table ? (
                            <TableContainer  component={Paper} className={classes.tableW}>
                                <Table stickyHeader className={classes.table} aria-label="simple table">
                                    <TableHead>
                                        <TableRow>
                                            {
                                                headers.map(val => {
                                                    return <TableCell> {val} </TableCell>
                                                })
                                            }
                                        </TableRow>
                                    </TableHead>
                                    {<TableBody>
                                        {body.map((rows, i) => (
                                            <TableRow key={i}>
                                                {headers.map((values) => {
                                                    return <TableCell> {rows[values]} </TableCell>
                                                })
                                                }
                                            </TableRow>
                                        ))}
                                    </TableBody>
                                    }
                                </Table>
                            </TableContainer>
                        ) : ""}

                    </Paper>
                </Grid>
            </Grid>
        </div>
    );
}


export { Quering };