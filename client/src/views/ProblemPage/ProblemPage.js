import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";

import Navbar from "components/Navbar/Navbar";
import CssBaseline from '@material-ui/core/CssBaseline';
import Container from '@material-ui/core/Container';
import MDEditor from '@uiw/react-md-editor';
import NavPills from "components/NavPills/NavPills.js";

import problemService from "services/problem.service";
import ProblemTable from "./Components/ProblemTable";
import ProblemSubmissions from "./Components/ProblemSubmission";

import CodeEditor from '@uiw/react-textarea-code-editor';
import EditorSettings from "views/ProblemPage/Components/EditorSettings";
import Button from "components/CustomButtons/Button.js";
import Footer from "components/Footer/Footer";
import { ToastContainer, toast } from 'react-toastify';

import evaluatorService from "services/evaluator.service";
import authenticationService from "services/authentication.service";
import Loading from "views/Loading";

const ProblemPage = () => {
    
    const [loading, setLoading] = useState(true);
    const [fetchingStatus, setFetchingStatus] = useState(200);
    const { problemName } = useParams();
    const [problem, setProblem] = useState({})

    const [code, setCode] = useState(`#include <stdio.h>\n\nint main() {\n  int a, b;\n\n  scanf("%d %d", &a, &b);\n  printf("%d", a + b);\n\n  return 0;\n}`);
    const [fontSize, setFontSize] = useState(18);
    const [lang, setLang] = useState("c");
    const [editorColor, setEditorColor] = useState("#e0e0e0");

    const editorStyles = () => {
        return {
            fontSize: fontSize,
            backgroundColor: editorColor,
            fontFamily: 'ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace',
        }
    }

    const fetchProblem = async() => {
        try {
            const res = await problemService.getByName(problemName);
            setProblem(res.data);
        } catch (err) {
          console.error(err);
        } finally {
            setLoading(false);
        }
    }

    // TODO submissions that have not been evaluated
    // should be shown on the ui with a "not evaluated yet" message 

    const handleCodeSubmission = async() => {
        try {
            await evaluatorService.createSubmission(code, lang, problem.id);
            toast.info("Submission Sent", {
                fontSize: 30,
                position: "bottom-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: false,
                draggable: false,
                progress: undefined,
            });
        } catch(err) {
            console.error(err);
        }
    }

    useEffect(fetchProblem, []);

    if (loading) {
        return <Loading/>
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
            <Navbar color="white" fixed ={false}/> 
            <CssBaseline/>
            <Container style={{marginTop: "90px", width: "90%"}} maxWidth={false} >
                <ProblemTable data={problem}/>
                <NavPills color = "info"
                        tabs = {[
                        {
                            tabButton: "Description",
                            tabContent: (
                                <div style={{border: "1px solid #bdbdbd", padding: "20px"}}>
                                    <MDEditor.Markdown style={{marginBottom: "20px"}} source={problem.description} />
                                </div>
                            )
                        },
                        {
                            tabButton: "Submissions",
                            tabContent: <ProblemSubmissions problem={problem}/>
                        },
                        ]}
                    />
                {authenticationService.isLoggedIn() &&
                <div style={{border: "1px solid #bdbdbd", padding: "20px", margin: "20px 0 0px 0"}}>
                    <EditorSettings 
                        setFontSize={setFontSize} 
                        setLang={setLang} 
                        setColor={setEditorColor}
                    /> 
                    <div style={{overflow: "auto", maxHeight: "500px"}}>
                        <CodeEditor 
                            value={code} 
                            language={lang}
                            minHeight={200}
                            placeholder="write/paste code in the editor"
                            onChange={(evn) => setCode(evn.target.value)}
                            padding={12}
                            style={editorStyles()}
                        /> 
                    </div>
                    <Button onClick={handleCodeSubmission} style={{marginTop: "12px", width: "100%"}} color="info">
                        Submit 
                    </Button> 
                </div>
                }
                <Footer/>
            </Container>
        </div>
    );
}

export default ProblemPage;