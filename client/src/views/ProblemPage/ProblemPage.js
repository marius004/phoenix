import React, {useEffect, useState} from "react";

import {useParams} from "react-router-dom";
import useMediaQuery from '@material-ui/core/useMediaQuery';
import useWindowDimensions from "hooks/window-dimensions.js";

import Navbar from "components/Navbar/Navbar";
import GridContainer from "components/Grid/GridContainer.js";
import GridItem from "components/Grid/GridItem";
import NavPills from "components/NavPills/NavPills.js";
import Button from "components/CustomButtons/Button.js";
import EditorSettings from "views/ProblemPage/Components/EditorSettings";

import CssBaseline from '@material-ui/core/CssBaseline';
import Container from '@material-ui/core/Container';
import CodeEditor from '@uiw/react-textarea-code-editor';

import ProblemTab from "views/ProblemPage/Components/ProblemTab";
import DescriptionTab from "views/ProblemPage/Components/DescriptionTab";
import SubmissionTab from "views/ProblemPage/Components/SubmissionTab";

import problemService from "services/problem.service";
import NotFound from "views/NotFound";
import Loading from "views/Loading";
import InternalServerError from "views/InternalServerError";

const ProblemPage = () => {

    const [loading, setLoading] = useState(true);
    const [fetchingStatus, setFetchingStatus] = useState(200);

    const matches = useMediaQuery('(min-width:960px)');
    const { height } = useWindowDimensions();
    const { problemName } = useParams();

    const [problem, setProblem] = useState({})

    const [code, setCode] = useState(`#include <iostream>\n\nusing namespace std;\n\nint main() {\n  int a, b;\n\n  cin >> a >> b;\n  cout << a + b << endl;\n\n  return 0;\n}`);
    const [fontSize, setFontSize] = useState(14);
    const [lang, setLang] = useState("cpp");
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
            if (err.message == "Network Error") {
                setFetchingStatus(500);
            } else {
                const status = err.response.status;
                setFetchingStatus(status);
            }
        } finally {
            setLoading(false);
        }
    }

    useEffect(fetchProblem, []);

    if (loading) {
        return <Loading/>
    }

    if (fetchingStatus === 404) {
        return <NotFound/>
    }

    if (fetchingStatus === 500) {
        return <InternalServerError/>
    }

    return (<div>
        <Navbar color="white" fixed ={false}/> 
        <CssBaseline/>
        <Container style={{marginTop: "85px"}} maxWidth={false} >
            <GridContainer>
                <GridItem xl={6} md={6} sm={12} style={!matches ? {} : {height: "89vh", overflow: "auto", borderRight: "2px solid #424242"}}>
                    <NavPills color = "warning"
                        tabs = {[
                        {
                            tabButton: "Enunt",
                            tabContent: <ProblemTab problem={problem}/>
                        },
                        {
                            tabButton: "Descriere",
                            tabContent: <DescriptionTab problem={problem}/>
                        },
                        {
                            tabButton: "Submisii",
                            tabContent: <SubmissionTab problem={problem}/>
                        },
                        ]}
                    />      
                </GridItem> 
                <GridItem xl={6} md={6} sm ={12} style ={!matches ? {} : { height: "89vh", overflowX: "auto" } } > 
                    { /* Add a line break if on mobile */ } 
                    {!matches && <hr/>} 
                    <EditorSettings 
                        setFontSize={setFontSize} 
                        setLang={setLang} 
                        setColor={setEditorColor}
                    /> 
                    <CodeEditor 
                        value={code} 
                        language={lang}
                        minHeight={250}
                        placeholder="Scrie cod aici!"
                        onChange={(evn) => setCode(evn.target.value)}
                        padding={12}
                        style={editorStyles()}
                    /> 
                    <Button style={{marginTop: "12px", width: "100%"}} color="info">
                        Submit 
                    </Button> 
                </GridItem>
            </GridContainer> 
        </Container> 
    </div>);
}

export default ProblemPage;