import React, {useEffect, useState} from "react";

import { makeStyles } from "@material-ui/core/styles";
import Slide from "@material-ui/core/Slide";
import IconButton from "@material-ui/core/IconButton";
import Dialog from "@material-ui/core/Dialog";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogContent from "@material-ui/core/DialogContent";
import Button from "components/CustomButtons/Button.js";
import Close from "@material-ui/icons/Close";
import styles from "assets/jss/material-kit-react/views/componentsSections/javascriptStyles.js";
import { useFormik } from 'formik';
import * as yup from 'yup';
import TextField from '@material-ui/core/TextField';
import authenticationService from "../../services/authentication.service.js";
import Alert from '@material-ui/lab/Alert';
import CloseIcon from '@material-ui/icons/Close';

const useStyles = makeStyles(styles);

const Transition = React.forwardRef(function Transition(props, ref) {
    return <Slide direction="up" ref={ref} {...props} />;
});
  
Transition.displayName = "Transition";

export default function SignupDialog(props) {
    const classes = useStyles();

    const [openErr, setOpenErr] = useState(false);
    const [signupErr, setSignupErr] = useState("");

    const validationSchema = yup.object({
        email: yup
            .string('Introduceti adresa de email')
            .email()
            .required('Adresa de email este necesara'),
        username: yup
          .string('Introduceți numele de utilizator')
          .required('Numele de utilizator este necesar'),
        password: yup
          .string('Introduceți parola')
          .min(8, 'Parola trebuie să aibă minim 8 caractere')
          .max(20, 'Parola trebuie să aibă maxim 20 caractere')
          .required('Parola este necesara'),
        confirmPassword: yup
          .string('Confirmati parola')
          .required('Confirmarea parolei este necesara')
          .oneOf([yup.ref('password'), null], 'Cele doua parole trebuie sa fie identice')
    });

    const clear = () => {
      formik.setErrors ({
        email: "",
        password: "",
        username: "",
        confirmPassword: "",
      });
      formik.setValues({
        email: "",
        password: "",
        username: "",
        confirmPassword: "",
      });
      formik.setTouched({
        email: false,
        password: false,
        username: false,
        confirmPassword: false,
      });

      setOpenErr(false);
      setSignupErr("");
    }

    useEffect(clear,[props.show]);

    const formik = useFormik({
        initialValues: {
            email: "",
            username: "",
            password: "",
            confirmPassword: "",
        },
        validationSchema: validationSchema,
        onSubmit: async (values) => {
          try {
            await authenticationService.signup(values.username, values.password,  values.email);
            clear();
            props.onSignup();
          } catch(err) {
            let message = err.response.data.message;
            setSignupErr(message);
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
        onClose={() => props.onSignup()}
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
            onClick={() => props.onSignup()}
          >
            <Close className={classes.modalClose} />
          </IconButton>
          <h3 className={classes.modalTitle} style={{textAlign: "center", fontWeight: "bold"}}>Creare cont</h3>
        </DialogTitle>
        <DialogContent style={{marginBottom: "0", paddingBottom: "0"}} 
          id="classic-modal-slide-description"
          className={classes.modalBody}>
             { openErr && signupErr !== "" && 
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
                {signupErr}
              </Alert>
            } 
            <form onSubmit={formik.handleSubmit}>
            <TextField
                style={{marginTop: "3px", padding: "2px"}}
                fullWidth
                id="email"
                name="email"
                label="Email"
                type="em"
                value={formik.values.email}
                onChange={formik.handleChange}
                error={formik.touched.email && Boolean(formik.errors.email)}
                helperText={formik.touched.email && formik.errors.email}
                />
                <TextField
                style={{marginTop: "3px", padding: "2px"}}
                fullWidth
                id="username"
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
                id="password"
                name="password"
                label="Password"
                type="password"
                value={formik.values.password}
                onChange={formik.handleChange}
                error={formik.touched.password && Boolean(formik.errors.password)}
                helperText={formik.touched.password && formik.errors.password}
                />
                <TextField
                style={{marginTop: "3px", padding: "2px"}}
                fullWidth
                id="confirmPassword"
                name="confirmPassword"
                label="Confirm password"
                type="password"
                value={formik.values.confirmPassword}
                onChange={formik.handleChange}
                error={formik.touched.confirmPassword && Boolean(formik.errors.confirmPassword)}
                helperText={formik.touched.confirmPassword && formik.errors.confirmPassword}
                />

                <Button style={{ width: "100%", marginTop: "30px" }} 
                        color="rose" round type="submit" 
                        onSubmit={formik.onSubmit}>
                    Inregistrare
                </Button>
            </form>
        </DialogContent>
      </Dialog>  
    );
}