import React, {useState, useEffect} from "react";
import Navbar from "components/Navbar/Navbar";
import { Container, CssBaseline } from "@material-ui/core";
import { useParams } from "react-router";
import { Link } from "react-router-dom";
import problemAPI from "api/problem";

import Loading from "views/Components/Loading";
import InternalServerError from "views/Components/InternalServerError";
import NotFound from "views/Components/NotFound";
import Unauthorized from "views/Components/Unauthorized";
import problemUtil from "util/problem";
import NavPills from "components/NavPills/NavPills.js";
import GeneralEditTab from "./EditTabs/GeneralEditTab";
import DescriptionEditTab from "./EditTabs/DescriptionEditTab";
import ShortDescriptionEditTab from "./EditTabs/ShortDescriptionEditTab";
import CreateTestEditTab from "./EditTabs/CreateTestEditTab";
import UpdateTests from "./EditTabs/UpdateTests";

export default function EditProblemPage() {
    const {problemName} = useParams();
    const [problem, setProblem] = useState({});
    const [loading, setLoading] = useState(true);
    const [fetchingStatus, setFetchingStatus] = useState(200);

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

    const problemLink = () => {
        return `/problems/${problemName}`;
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

    if (!problemUtil.canUserEditProblem(problem)) {
        return <Unauthorized/>
    }

    return (
    <div>
        <CssBaseline/>
        <Navbar color="white" fixed ={false}/> 
        <Container style={{border: "1px solid grey", marginTop: "100px", marginBottom: "50px"}}>
            <h3>
                Edit Problem  {"  "}
                <Link to={problemLink} style={{color: "blue"}}>{problemName}</Link>
            </h3>
            <NavPills color = "info"
                    tabs = {[
                    {
                        tabButton: "General",
                        tabContent: <GeneralEditTab problem={problem} setProblem={setProblem}/>,
                    },
                    {

                        tabButton: "Description",
                        tabContent: <DescriptionEditTab problem={problem} setProblem={setProblem}/>,
                    },
                    {
                        tabButton: "Short Description",
                        tabContent: <ShortDescriptionEditTab problem={problem} setProblem={setProblem}/>,
                    }, 
                    {
                        tabButton: "Create Test",
                        tabContent: <CreateTestEditTab problem={problem} />,
                    }, 
                    {
                        tabButton: "Update Tests",
                        tabContent: <UpdateTests problem={problem} />,
                    }
                ]}
            />
        </Container>
    </div>
  );
}
