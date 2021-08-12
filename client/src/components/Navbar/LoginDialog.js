import React, { useEffect, useState } from "react";

import { makeStyles } from "@material-ui/core/styles";
import Slide from "@material-ui/core/Slide";
import IconButton from "@material-ui/core/IconButton";
import Dialog from "@material-ui/core/Dialog";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogContent from "@material-ui/core/DialogContent";
import Button from "components/CustomButtons/Button.js";
import Close from "@material-ui/icons/Close";
import styles from "assets/jss/material-kit-react/views/componentsSections/javascriptStyles.js";
import {useFormik} from 'formik';
import * as yup from 'yup';
import TextField from '@material-ui/core/TextField';
import authenticationAPI from "api/authentication";
import Alert from '@material-ui/lab/Alert';
import CloseIcon from '@material-ui/icons/Close';

const useStyles = makeStyles(styles);

const validationSchema = yup.object({
    username: yup
      .string('Introduceți numele de utilizator')
      .required('Numele de utilizator este necesar'),
    password: yup
      .string('Introduceți parola')
      .min(8, 'Parola trebuie să aibă minim 8 caractere')
      .max(20, 'Parola trebuie să aibă maxim 20 caractere')
      .required('Parola este necesara'),
});

const Transition = React.forwardRef(function Transition(props, ref) {
    return <Slide direction="up" ref={ref} {...props} />;
});
  
Transition.displayName = "Transition";

export default function LoginDialog(props) {
    const classes = useStyles();
    const [loginErr, setLoginErr] = useState("");
    const [openErr, setOpenErr] = useState(false);

    const clear = () => {
      formik.setErrors ({
        password: "",
        username: "",
      });
      formik.setValues({
        password: "",
        username: "",
      });
      formik.setTouched({
        password: false,
        username: false,
      });

      setLoginErr("");
      setOpenErr(false);
    }

    useEffect(clear, [props.show]);

    const formik = useFormik({
        initialValues: {
            username: "",
            password: "",
        },
        validationSchema: validationSchema,
        onSubmit: async (values) => {
          try {
            const res = await authenticationAPI.login(values.username, values.password);
            console.log(res);
            window.location.reload()
          } catch(err) {
            console.error(err);
            let message = err.response.data.message;
            setLoginErr(message);
            setOpenErr(true);
          }
        },
      });

    return (
        <Dialog
        classes={{
          root: classes.center,
          paper: classes.modal,
        }}
        open={props.show}
        TransitionComponent={Transition}
        keepMounted
        onClose={() => props.onLogin()}
        aria-labelledby="classic-modal-slide-title"
        aria-describedby="classic-modal-slide-description"
      >
        <DialogTitle
          id="classic-modal-slide-title"
          disableTypography
          className={classes.modalHeader}
        >
          <IconButton
            className={classes.modalCloseButton}
            key="close"
            aria-label="Close"
            color="inherit"
            onClick={() => props.onLogin()}
          >
            <Close className={classes.modalClose} />
          </IconButton>
          <h3 className={classes.modalTitle} style={{textAlign: "center", fontWeight: "bold"}}>Autentificare</h3>
        </DialogTitle>
        <DialogContent style={{marginBottom: "0", paddingBottom: "0"}} 
          id="classic-modal-slide-description"
          className={classes.modalBody}>
            { openErr && loginErr !== "" && 
             <Alert
             severity="error"
              action={
                <IconButton
                  aria-label="close"
                  color="inherit"
                  size="small"
                  onClick={() => {
                    setOpenErr(false);
                  }}
                >
                <CloseIcon fontSize="inherit" />
                </IconButton>
              }>
                {loginErr}
              </Alert>
            } 

            <form onSubmit={formik.handleSubmit}>
                <TextField
                style={{marginTop: "3px", padding: "2px"}}
                fullWidth
                id="login-username"
                name="username"
                label="Username"
                value={formik.values.username}
                onChange={formik.handleChange}
                error={formik.touched.username && Boolean(formik.errors.username)}
                helperText={formik.touched.username && formik.errors.username}
                />
                <TextField
                style={{marginTop: "3px", padding: "2px"}}
                fullWidth
                id="login-password"
                name="password"
                label="Password"
                type="password"
                value={formik.values.password}
                onChange={formik.handleChange}
                error={formik.touched.password && Boolean(formik.errors.password)}
                helperText={formik.touched.password && formik.errors.password}
                />

                <Button style={{ width: "100%", marginTop: "30px" }} 
                    color="rose" round type="submit">
                    Autentificare
                </Button>
            </form>
        </DialogContent>
      </Dialog>  
    );
}