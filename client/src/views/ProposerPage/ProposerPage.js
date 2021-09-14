import React, {useState} from 'react';
import { Container, CssBaseline } from '@material-ui/core';
import Navbar from 'components/Navbar/Navbar';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import problemAPI from 'api/problem';
import { ToastContainer, toast } from 'react-toastify';

const useStyles = makeStyles((theme) => ({
    root: {
      '& > *': {
        margin: theme.spacing(1),
        width: '25ch',
      },
    },
}));

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

export default function ProposerPage() {
   const classes = useStyles();
   const [name, setName] = useState("");

   const handleSubmit = async(e) => {
        e.preventDefault();

        const problem = {
           name: name,
           memoryLimit: 65536,
           stackLimit: 16384,
           timeLimit: 1,
           difficulty: "easy",
           stream: "console",
           grade: "9",
        };

        try {
            await problemAPI.create(problem);
            toast.success("Problem created successfully!", toastConfig);
        } catch(err) {
            console.error(err);
            const message = err.response.data.message;
            toast.error(message, toastConfig);
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
        <CssBaseline/>
        <Navbar color="white" fixed ={false} />
        <Container style={{marginTop: "100px"}}>
            <h2 style={{textAlign: "center"}}>Panou Propunator</h2>
            <form action="" onSubmit={handleSubmit} style={{border: "1px solid #9e9e9e", padding: "4px"}} className={classes.root} noValidate autoComplete="off">
                <h3>Creare Problema</h3>
                <TextField
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    id="outlined-secondary"
                    label="Nume problema"
                    variant="outlined"
                    color="secondary"
                    style={{width: "99%"}}
                />
                <br/>
               <Button type="submit" variant="contained" color="primary" >
                    Creare
                </Button>
            </form>
        </Container>
       </div>
   );
}