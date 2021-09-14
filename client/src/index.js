import React from "react";
import ReactDOM from "react-dom";
import { createBrowserHistory } from "history";
import { Router, Route, Switch } from "react-router-dom";
import "assets/scss/material-kit-react.scss";
import ProblemPage from "views/ProblemPage/ProblemPage.js";
import HomePage from "views/HomePage/HomePage.js";
import NotFound from "views/Components/NotFound";
import 'react-toastify/dist/ReactToastify.css';
import ProfilePage from "views/ProfilePage/ProfilePage";
import SubmissionsPage from "views/SubmissionsPage/SubmissionsPage";
import SubmissionPage from "views/SubmissionPage/SubmissionPage";
import ProblemsPage from "views/ProblemsPage/ProblemsPage";
import RequireProposer from "components/Authorization/RequireProposer";
import Unauthorized from "views/Components/Unauthorized";
import ProposerPage from "views/ProposerPage/ProposerPage";
import AdminPage from "views/AdminPage/AdminPage";
import RequireAdmin from "components/Authorization/RequireAdmin";
import EditProblemPage from "views/ProblemPage/EditProblemPage";

const hist = createBrowserHistory();

ReactDOM.render(
  <Router history={hist}>
    <Switch>
      <Route exact path="/" component={HomePage} />
      <Route exact path="/problems/:problemName" component={ProblemPage} />
      <Route exact path="/problems/:problemName/edit" component={RequireProposer(Unauthorized, EditProblemPage)} />
      <Route exact path="/problems" component={ProblemsPage}/>
      <Route path="/profile/:username" component={ProfilePage} />
      <Route path="/submissions/:submissionId" component={SubmissionPage} />
      <Route path="/submissions" component={SubmissionsPage} />
      <Route path="/proposer" component={RequireProposer(Unauthorized, ProposerPage)} />
      <Route path="/admin" component={RequireAdmin(Unauthorized, AdminPage)} />
      <Route path="*"component={NotFound} />
    </Switch>
  </Router>,
  document.getElementById("root")
);