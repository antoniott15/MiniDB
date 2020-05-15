import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            flexGrow: 1,
        },
        menuButton: {
            marginRight: theme.spacing(2),
        },
        typo: {
            justifyContent: "center",
        }
    }),
);

const NavBar = () => {
    const classes = useStyles();

    return (
        <div className={classes.root}>
            <AppBar position="static" >
                <Toolbar variant="dense" className={classes.typo}>
                    <Typography variant="h6" color="secondary">
                        Base de datos 2
                    </Typography>
                </Toolbar>
            </AppBar>
        </div>
    );
}

export { NavBar };