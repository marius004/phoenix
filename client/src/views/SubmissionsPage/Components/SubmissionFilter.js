import React, {useState} from "react";
import PropTypes from "prop-types";
import TextField from '@material-ui/core/TextField';
import { makeStyles } from '@material-ui/core/styles';
import { Button, MenuItem} from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
    root: {
      '& .MuiTextField-root': {
        margin: theme.spacing(1),
        width: '25ch',
      },
    },
}));
  
export default function SubmissionFilter(props) {
    const classes = useStyles();

    const textFieldStyles = () => {
      return {width: "100%", margin: "0 auto", marginBottom: "12px", textAlign: "left"}
    };

    const handlePageChange = (e) => {
      const value = e.target.value;

      if (value == "")
        props.setPage("");
      else {
        const nr = parseInt(value);

        if(nr >= 0)
          props.setPage(value);
      }
    }

    const handleScoreChange = (e) => {
      const value = e.target.value;

      if (value === "")
        props.setScore("");
      else {
        const nr = parseInt(value);

        if(nr >= -1 && nr <= 100)
          props.setScore(value);
      }
    }

    return (
    <div style={{textAlign: "center"}}>
    <form className={classes.root} noValidate onSubmit={props.onSubmit}>
      <div>
        <TextField
          style={textFieldStyles()}
          id="submission-form-page"
          label="Page"
          type="number"
          value={props.page}
          onChange={handlePageChange}
          variant="outlined"
        />
        <TextField
          style={textFieldStyles()}
          id="submission-form-username"
          label="Username"
          type="text"
          value={props.username}
          onChange={(e) => props.setUsername(e.target.value)}
          variant="outlined"
        />
        <TextField
          style={textFieldStyles()}
          id="submission-form-problem"
          label="Problem"
          type="text"
          value={props.problemName}
          onChange={(e) => props.setProblemName(e.target.value)}
          variant="outlined"
        />
        <TextField 
            style={textFieldStyles()}
            variant="outlined"
            id="" 
            label="Status" 
            value={props.status}
            onChange={(e) => props.setStatus(e.target.value)}
            select>
                <MenuItem value="-">-</MenuItem>
                <MenuItem value="waiting">Waiting for evaluation</MenuItem>
                <MenuItem value="working">Evaluating</MenuItem>
                <MenuItem value="finished">Evaluated</MenuItem>
        </TextField>
        <TextField 
            style={textFieldStyles()}
            variant="outlined"
            id="" 
            label="Language" 
            value={props.language}
            onChange={(e) => props.setLanguage(e.target.value)}
            select>
                <MenuItem value="-">-</MenuItem>
                <MenuItem value="c">C</MenuItem>
        </TextField>
        <TextField
          style={textFieldStyles()}
          id="submission-form-score"
          label="Score"
          type="number"
          value={props.score}
          onChange={handleScoreChange}
          variant="outlined"
        />
       <Button 
            type="submit"
            variant="contained" 
            color="primary" 
            style={{width: "100%"}}
          >
           Search
        </Button>
      </div>
    </form>
    </div>
    );
}

SubmissionFilter.propTypes = {
    setUsername: PropTypes.func.isRequired,
    setProblemName: PropTypes.func.isRequired,
    setStatus: PropTypes.func.isRequired,
    setLanguage: PropTypes.func.isRequired,
    setScore: PropTypes.func.isRequired,
    setPage: PropTypes.func.isRequired,

    username: PropTypes.string.isRequired,
    problemName: PropTypes.string.isRequired,
    status: PropTypes.string.isRequired,
    language: PropTypes.string.isRequired,
    score: PropTypes.string.isRequired,
    page: PropTypes.string.isRequired,

    onSubmit: PropTypes.func.isRequired
};