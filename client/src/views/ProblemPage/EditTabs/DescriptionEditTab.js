import React, {useState} from "react";
import MDEditor from '@uiw/react-md-editor';
import { Button } from "@material-ui/core";
import problemAPI from "api/problem";
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

export default function DescriptionEditTab({problem, setProblem}) {
    const [description, setDescription] = useState(problem.description);

    const handleDescriptionUpdate = async() => {
        try {
            await problemAPI.update(problem.name, {
                description,
            });

            setProblem({
                ...problem, 
                description,
            });

            toast.success("Problem description updated successfully!", toastConfig);
        } catch(err) {
            console.error(err);
            const message = err?.response?.data?.message;

            if (message != null)
                toast.error(`Could not update problem description: ${message}`, toastConfig);
            else
                toast.error("Could not update problem description", toastConfig);
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
            <div style={{border: "1px solid #01579b", padding: "12px", marginBottom: "20px"}}>
                <h4>
                    The problem description must be written in 
                    <a href="https://www.markdownguide.org/"> markdown </a> !
                </h4>
                
                <h4 style={{fontWeight: "bold"}}>Edit problem description</h4>
                <MDEditor
                    value={description}
                    onChange={setDescription}
                    style={{marginBottom: "12px"}}
                />

                <Button onClick={handleDescriptionUpdate} type="submit" variant="contained" color="primary" >
                    Update
                </Button>
            </div>
        </div>
    );
}