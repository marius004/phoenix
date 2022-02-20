import { Select } from '@material-ui/core';
import testAPI from 'api/problem-test';
import React, {useEffect, useState} from 'react';
import { MenuItem } from '@material-ui/core';
import { InputLabel } from '@material-ui/core';
import { FormControl } from '@material-ui/core';
import { Box } from '@material-ui/core';
import { Button } from '@material-ui/core';
import { ToastContainer, toast } from 'react-toastify';

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

export default function UpdateTests({problem}) {
    const [tests, setTests] = useState([]);

    const [testId, setTestId] = useState(-1);
    const [test, setTest] = useState({});

    const fetchProblemTests = async() => {
        const tests = await testAPI.getProblemTests(problem.name);
        setTests(tests);
    };

    const handleTestIdChange = async() => { 
        if(testId === -1) {
            setTest({});
            return;
        }

        const test = await testAPI.getProblemTestById(problem.name, testId);
        setTest(test);
    };

    const handleTestUpdate = async() => {
        try {
            await testAPI.updateProblemTestById(problem.name, testId, test.test.score, test.input, test.output);
            toast.success("Test updated successfully!", toastConfig);
        } catch(err) {
            console.error(err);
            const message = err?.response?.data?.message;

            if (message != null)
                toast.error(`Test update failed: ${message}`, toastConfig);
            else
                toast.error("Test update failed", toastConfig);
        }
    };

    const removeTest = (testId) => {
        let list = [...tests];
        list = list.filter(test => test.id != testId);
        setTests(list);
    };

    const handleDeleteTest = async() => {
        if (confirm("Are you sure you want to delete this test?")) {
            try {
                await testAPI.deleteProblemTestById(problem.name, testId);
                toast.success(`Test ${testId} deleted!`, toastConfig);
                removeTest(testId);
                setTestId(-1);
            } catch(err) {
                console.error(err);
                const message = err?.response?.data?.message;

                if (message != null)
                    toast.error(`Could not delete test: ${message}`, toastConfig);
                else
                    toast.error("Could not delete test", toastConfig);
            }
        }
    };

    const setTestInput = (e) => {
        const value = e.target.value;

        setTest({
            ...test,
            input: value,
        });
    }

    const setTestScore = (e) => {
        const value = e.target.value;

        setTest({
            ...test,
            test: {
                score: score,
            }
        });
    }

    const setTestOutput = (e) => {
        const value = e.target.value;

        setTest({
            ...test,
            output: value,
        });
    }

    useEffect(fetchProblemTests, []);
    useEffect(handleTestIdChange, [testId]);

    return (
        <div style={{marginBottom: "20px"}}>
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
            <Box sx={{ minWidth: 120 }}>
                <FormControl fullWidth>
                    <InputLabel id="demo-simple-select-label">Test Id</InputLabel>
                    <Select
                        labelId="test-id"
                        id="test-id"
                        value={testId}
                        label="test-id"
                        onChange={(e) => setTestId(e.target.value)}
                    >
                        <MenuItem value={-1}>-</MenuItem>
                        {
                            tests.map((test) => {
                            return <MenuItem key={test.id} value={test.id}>{test.id}</MenuItem>
                            })
                        }
                    </Select>
                </FormControl>
            </Box>

            {testId !== -1 &&
                <>
                {/* TODO score */}
                    {/* <h4>Score: </h4>
                    <input
                        value={test.test.score}
                        onChange={setTestScore}
                        style={{resize: "none", width: "100%", minHeight: "200px"}}>
                    </input> */}

                    <h4>Input: </h4>
                    <textarea 
                        value={test.input}
                        onChange={setTestInput}
                        style={{resize: "none", width: "100%", minHeight: "200px"}}
                    />

                    <h4>Output: </h4>
                    <textarea 
                        value={test.output}
                        onChange={setTestOutput}
                        style={{resize: "none", width: "100%", minHeight: "200px"}}
                    />

                    <Button onClick={handleTestUpdate} type="submit" variant="contained" color="primary" style={{marginRight: "2px"}} >
                        Update
                    </Button>

                    <Button onClick={handleDeleteTest} type="submit" variant="contained" style={{backgroundColor: "#d50000", color: "white"}}>
                        Delete
                    </Button>
                </>
            }
        </div>
    );
}