import React, {useEffect, useState} from "react";
import {Link, useParams} from "react-router-dom";

import Navbar from "components/Navbar/Navbar";
import CssBaseline from '@material-ui/core/CssBaseline';
import Container from '@material-ui/core/Container';
import MDEditor from '@uiw/react-md-editor';
import NavPills from "components/NavPills/NavPills.js";

import problemAPI from "api/problem";
import ProblemTable from "./Components/ProblemTable";
import ProblemSubmissions from "./Components/ProblemSubmission";

import CodeEditor from '@uiw/react-textarea-code-editor';
import EditorSettings from "views/ProblemPage/Components/EditorSettings";
import Button from "components/CustomButtons/Button.js";
import Footer from "components/Footer/Footer";
import { ToastContainer, toast } from 'react-toastify';

import submissionAPI from "api/submission";
import authenticationUtil from "util/authentication";

import Loading from "views/Components/Loading";
import InternalServerError from "views/Components/InternalServerError";
import NotFound from "views/Components/NotFound";
import problemUtil from "util/problem";

const toastConfig = {
    fontSize: 30,
    position: "bottom-right",
    autoClose: 5000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: false,
    draggable: false,
    progress: undefined,
}

const ProblemPage = () => {
    
    const [loading, setLoading] = useState(true);
    const [submissionId, setSubmissionId] = useState(-1);
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
            const res = await problemAPI.getByName(problemName);
            setProblem(res);
        } catch (err) {
            console.error(err);

            if (err.message == "Network Error") {
                setFetchingStatus(500)
                return;
            }

            if (err.response.status === 404) {
                setFetchingStatus(404);
                return;
            }

            if (err?.response?.status) {
                setFetchingStatus(err.response.status);
                return;
            }

        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        if (submissionId < 0)
            return;

        const timer = setInterval(async() => {
            console.log("iterating...")
            
            if (submissionId < 0)
                return;

            const submission = await submissionAPI.getById(submissionId);

            if (submission.status === "finished") {

                if (submission.score === 100) {
                    toast.success(<p>Submission #{submissionId} evaluated <br/> Score: {submission.score}</p>, toastConfig);
                } else if(submission.score >= 50) {
                    toast.warning(<p>Submission #{submissionId} evaluated <br/> Score: {submission.score}</p>, toastConfig);
                } else {
                    toast.error(<p>Submission #{submissionId} evaluated <br/> Score: {submission.score}</p>, toastConfig);
                }

                clearInterval(timer);
                setSubmissionId(-1);
            }
            
        }, 2000) 

        return () => {
            clearInterval(timer);
        }
    }, [submissionId])
       
    const handleCodeSubmission = async() => {
        try {
            const res = await submissionAPI.create(code, lang, problem.id);
            setSubmissionId(res.id);
            
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

    const editProblemLink = () => {
        return `/problems/${problem.name}/edit`;
    }

    useEffect(fetchProblem, []);
    
    if (loading) {
        return <Loading/>
    }

    // 4xx (bad request or not found)
    if (Math.round(fetchingStatus / 100) == 4) {
        return <NotFound message="Problema cautata nu a fost gasita"/>   
    }

    // 5xx (internal server error)
    if (Math.round(fetchingStatus / 100) === 5) {
        return <InternalServerError/>
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
                                <div style={{border: "1px solid #bdbdbd", padding: "12px"}}>
                                    {problemUtil.canUserEditProblem(problem) &&
                                        <h3>
                                            Problema {problem.name} {"  "}
                                            <a style={{color: "blue"}} href={editProblemLink()}>[EDIT]</a>
                                        </h3>
                                    }
                                    {!problemUtil.canUserEditProblem(problem) &&
                                        <h3>Problema {problem.name}</h3>
                                    }
                                    <MDEditor.Markdown style={{marginBottom: "12px"}} source={problem.description} />
                                </div>
                            )
                        },
                        {
                            tabButton: "Submissions",
                            tabContent: <ProblemSubmissions problem={problem}/>
                        },
                        ]}
                    />
                {authenticationUtil.isUserLoggedIn() &&
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