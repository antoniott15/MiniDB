import React, { useEffect, useState } from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import { allTables, selectToTable } from "../helpers/apicall";
import Button from '@material-ui/core/Button';
import { Select } from "../Quering/index";
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

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
        buttons: {
            '& > *': {
                margin: theme.spacing(1),
            },
        },
        table: {
            minWidth: 500,
        },
        tableW:{
            maxHeight: 440,
        }
    }),
);

export const Tables = () => {
    const classes = useStyles();
    const [tables, setTables] = useState<Array<String>>([])
    const [headers, setHeaders] = useState<string[]>([]);
    const [body, setBody] = useState<any[]>([]);

    useEffect(() => {
        allTables().then((t) => {
            const tables = t.data.data;
            const newTables = new Array<String>();
            for (let elements of tables) {
                const value = String(elements);
                const tables = value.replace("Tables/", "");
                console.log(tables)
                newTables.push(tables)
            }
            setTables(newTables);
        }).catch((err) => {
            console.log(err)
        });
    }, [])


    const TableClicked = (e: any) => {
        const selected: Select = {
            table: e,
            query: "select",
            attribs: ["*"]
        }

        selectToTable(selected).then((res) => {
            setHeaders([]);
            setBody([]);
            const data = res.data.data;
            if(data){
                setHeaders(data[0].Headers);
                var arr = new Array<any>();
                for (const elements of data) {
                    arr.push(elements.Attribs)
                }
                setBody(arr);
            }
        })

    }

    return (
        <div className={classes.root}>
            <Grid container spacing={3}>
                <Grid item xs={3}>
                    <Paper className={classes.paper}>
                        <div className={classes.buttons}>
                            {
                                tables.map((val) => {
                                    return <Button color="secondary" size="large" onClick={() => TableClicked(val)}>{val.replace("./","")}</Button>
                                })
                            }
                        </div>
                    </Paper>
                </Grid>
                <Grid item xs={6} sm={9}>
                    <Paper className={classes.paper}>
                        <TableContainer component={Paper} className={classes.tableW}>
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
                    </Paper>
                </Grid>
            </Grid>
        </div>
    );
}