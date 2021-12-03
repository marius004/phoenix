import React, {useState} from "react";
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import problemAPI from "api/problem";
import { ToastContainer, toast } from 'react-toastify';
import { useHistory } from "react-router";

const useStyles = makeStyles((theme) => ({
    root: {
      '& > *': {
        margin: theme.spacing(1),
        width: '25ch',
      },
    },
}));

const toastConfig = {
    fontSize: 30,
    position: "bottom-right",
    autoClose: 5000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: false,
    draggable: false,
    progress: undefined,
};

export default function GeneralEditTab({problem, setProblem}) {
    const classes = useStyles();
    const history = useHistory();

    const [memoryLimit, setMemoryLimit] = useState(problem.memoryLimit);
    const [stackLimit, setStackLimit] = useState(problem.stackLimit);
    const [timeLimit, setTimeLimit] = useState(problem.timeLimit);
    const [grade, setGrade] = useState(problem.grade);

    const onGradeChange = (e) => {
        const value = e.target.value;
        if (value >= 9 && value <= 11)
            setGrade(value);
    }

    const onMemoryLimitChange = (e) => {
        const value = e.target.value;
        if (value >= 0 && value <= 131072)
            setMemoryLimit(value);
    }

    const onStackLimitChange = (e) => {
        const value = e.target.value;
        if (value >= 0 && value <= 131072)
            setStackLimit(value);
    }

    const onTimeLimitChange = (e) => {
        const value = e.target.value;
        if (value >= 0 && value <= 10)
            setTimeLimit(value);
    }

    const handleProblemUpdate = async() => {
        try {
            await problemAPI.update(problem.name, {
                grade: grade,
                timeLimit: parseFloat(timeLimit),
                memoryLimit: parseFloat(memoryLimit),
                stackLimit: parseFloat(stackLimit), 
            });

            setProblem({
                ...problem, 
                grade: grade, 
                timeLimit: timeLimit, 
                memoryLimit: memoryLimit,
                stackLimit: stackLimit,
            });

            toast.success("Problem updated successfully!", toastConfig);
        } catch(err) {
            console.error(err);
            const message = err?.response?.data?.message;

            if (message != null)
                toast.error(`Could not update problem: ${message}`, toastConfig);
            else
                toast.error("Could not update problem", toastConfig);
        }
    };

    const handleDeleteProblem = async () => {
        if (confirm("Are you sure you want to delete this problem?")) {
            try {
                await problemAPI.delete(problem.name);
                toast.success(`Problem ${problem.name} deleted! Redirecting you to the home page!`, toastConfig);
                
                setTimeout(() => {
                    history.push("/");
                }, 2000);
            } catch(err) {
                console.error(err);
                const message = err?.response?.data?.message;

                if (message != null)
                    toast.error(`Could not update problem: ${message}`, toastConfig);
                else
                    toast.error("Could not update problem", toastConfig);
            }
        }
    };

    return (
        <div>
            <ToastContainer
                position="bottom-right"
                autoClose={5000}
                hideProgressBar={false}
                newestOnTop={false}
                closeOnClick
                rtl={false}
                pauseOnFocusLoss={false}
                draggable={false}
                pauseOnHover={false}
            />
            <div style={{border: "1px solid #01579b", padding: "4px", marginBottom: "20px"}} className={classes.root}>
                <TextField
                    value={grade}
                    onChange={onGradeChange}
                    type="number"
                    id="outlined-secondary"
                    label="Grade"
                    variant="outlined"
                    color="secondary"
                    style={{width: "99%"}}
                />
                <TextField
                    value={memoryLimit}
                    onChange={onMemoryLimitChange}
                    type="number"
                    id="outlined-secondary"
                    label="Memory Limit(KB)"
                    variant="outlined"
                    color="secondary"
                    style={{width: "99%"}}
                />
                <TextField
                    value={stackLimit}
                    onChange={onStackLimitChange}
                    type="number"
                    id="outlined-secondary"
                    label="Stack Limit(KB)"
                    variant="outlined"
                    color="secondary"
                    style={{width: "99%"}}
                />
                <TextField
                    value={timeLimit}
                    onChange={onTimeLimitChange}
                    type="number"
                    id="outlined-secondary"
                    label="Time Limit(seconds)"
                    variant="outlined"
                    color="secondary"
                    style={{width: "99%"}}
                />
                <br/>
                <Button onClick={handleProblemUpdate} type="submit" variant="contained" color="primary" >
                    Update
                </Button>
                <Button onClick={handleDeleteProblem} type="submit" variant="contained" style={{backgroundColor: "#d50000", color: "white"}} >
                    Delete Problem
                </Button>
            </div>
        </div>
    );
}