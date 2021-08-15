import React, {useState, useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardMedia from '@material-ui/core/CardMedia';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import Collapse from '@material-ui/core/Collapse';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import { cyan } from '@material-ui/core/colors';
import FavoriteIcon from '@material-ui/icons/Favorite';
import ShareIcon from '@material-ui/icons/Share';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import MoreVertIcon from '@material-ui/icons/MoreVert';
import MDEditor from '@uiw/react-md-editor';
import { CssBaseline } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {},
  media: {
    height: 0,
    paddingTop: '56.25%', // 16:9
  },
  expand: {
    transform: 'rotate(0deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
      duration: theme.transitions.duration.shortest,
    }),
  },
  expandOpen: {
    transform: 'rotate(180deg)',
  },
  avatar: {
    backgroundColor: cyan[600],
  },
}));

export default function ProblemCard({problem, submission}) {
  const classes = useStyles();
  const [expanded, setExpanded] = useState(false);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };

  const redirectToProblemPage = () => {
      window.location.href = `/problems/${problem.name}`;
  }

  return (
    <div>
        <h3>Problem Description</h3>
        <Card className={classes.root}>
        <CardHeader
            avatar={
            <Avatar aria-label="recipe" className={classes.avatar}>
                PB
            </Avatar>
            }
            title={problem.name ? problem.name : "title"}
            subheader={problem.stream ? problem.stream : "console"}
        />
        <CardContent>
            <Typography variant="body2" color="textSecondary" component="p">
                <MDEditor.Markdown style={{marginBottom: "20px"}} source={problem.shortDescription} />
            </Typography>
        </CardContent>
        <CardActions disableSpacing>
            <IconButton aria-label="share">
            <ShareIcon onClick={redirectToProblemPage}/>
            </IconButton>
            <IconButton
            className={clsx(classes.expand, {
                [classes.expandOpen]: expanded,
            })}
            onClick={handleExpandClick}
            aria-expanded={expanded}
            aria-label="show more"
            >
            <ExpandMoreIcon />
            </IconButton>
        </CardActions>
        <Collapse in={expanded} timeout="auto" unmountOnExit>
            <CardContent>
              <Typography>Grade: {problem.grade}</Typography>
              <Typography>Difficulty: {problem.difficulty}</Typography>
              <Typography paragraph>Stream: {problem.stream}</Typography>
              <Typography>Time Limit: {problem.timeLimit}s</Typography>
              <Typography>Memory Limit: {problem.memoryLimit} KB</Typography>
              <Typography paragraph>Stack Limit: {problem.stackLimit} KB</Typography>
            </CardContent>
        </Collapse>
        </Card> 
    </div>
  );
}
