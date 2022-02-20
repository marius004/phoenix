import React, {useState} from "react";
import { TextField } from "@material-ui/core";
import { Button } from "@material-ui/core";
import testAPI from "api/problem-test";
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

export default function CreateTestEditTab({problem}) {

    const [score, setScore] = useState(10);
    const [input, setInput] = useState("");
    const [output, setOutput] = useState("");

    const onScoreChange = (e) => {
        const value = e.target.value;

        if (value > 0 && value <= 100)
            setScore(value);
    }

    const handleCreateTest = async() => {
        if (input.trim() === "" && output.trim() === "") {
            toast.warning("Please enter a valid input and output", toastConfig);
            return;
        }
        try {
            await testAPI.createProblemTest(problem.name, score, input, output);
            toast.success("Test created successfully!", toastConfig);
        } catch(err) {
            console.error(err);
            const message = err?.response?.data?.message;

            if (message != null)
                toast.error(`Could not create test: ${message}`, toastConfig);
            else
                toast.error("Could not create test", toastConfig);
        }
    }

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
            <div style={{border: "1px solid grey", padding: "12px", marginBottom: "20px"}}>
                <TextField
                    value={score}
                    onChange={onScoreChange}
                    type="number"
                    id="outlined-secondary"
                    label="Score"
                    variant="outlined"
                    color="secondary"
                    style={{width: "99%", marginBottom: "10px"}}
                />
                
                <h4>Input: </h4>
                <textarea 
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    style={{resize: "none", width: "100%", minHeight: "200px"}}
                />

                <h4>Output: </h4>
                <textarea 
                    value={output}
                    onChange={(e) => setOutput(e.target.value)}
                    style={{resize: "none", width: "100%", minHeight: "200px"}}
                />
                
                <Button onClick={handleCreateTest} variant="contained" color="primary" >
                    Create
                </Button>
            </div>
        </div>
    );
} 