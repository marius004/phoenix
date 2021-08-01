import React from "react";
import { makeStyles } from '@material-ui/core/styles';
import CircularProgress from '@material-ui/core/CircularProgress';

const useStyles = makeStyles((theme) => ({
    root: {
      display: 'block',
      '& > * + *': {
        marginLeft: theme.spacing(5),
      },
    },
  }));

const Loading = () => {
    const classes = useStyles();

    return (<>
        <div className={classes.root} style={{ position: "absolute", top: "50%", left: "50%", transform: "translate(-50%,-50%)" }}>
            <CircularProgress />
        </div>
    </>);
};

export default Loading;
