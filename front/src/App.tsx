import React from 'react';
import { Quering } from "./components/Quering/index"
import { NavBar } from "./components/organism/navbar"
import { Tables } from "./components/tables/index"
import { makeStyles, createStyles, Theme, createMuiTheme, MuiThemeProvider } from '@material-ui/core/styles';
import { Grid } from '@material-ui/core';

//import { MuiThemeProvider, createMuiTheme } from '@material-ui/styles';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      margin: 0,
      top: 0,
    },
  }),
);

const THEME = createMuiTheme({
  typography: {
    "fontFamily": `"Quicksand", "Roboto", "Arial", sans-serif`,
    "fontSize": 14,
    "fontWeightLight": 300,
    "fontWeightRegular": 400,
    "fontWeightMedium": 500
  },
  palette: {
    primary: {
      light: '#BDFFFD',
      main: '#9FFFF5',
      dark: '#BDFFFD',
      contrastText: '#fff',
    },
    secondary: {
      light: '##79ed7e',
      main: '#017a87',
      dark: '#004e5a',
      contrastText: '#000',
    },
    error: {
      main: '#e7504d'
    }
  }
});

function App() {
  const classes = useStyles();
  return (
    <MuiThemeProvider theme={THEME}>
      <div className={classes.root}>
          <NavBar></NavBar>
          <Grid>
            <Quering></Quering>
            <Tables></Tables>
          </Grid>
      </div>
    </MuiThemeProvider>
  );
}

export default App;


