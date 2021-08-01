import Badge from 'components/Badge/Badge.js';

const DifficultyBadge = ({ difficulty }) => {
  switch(difficulty) {
      case "contest":
          return <Badge color="rose">concurs</Badge>;
      case "hard":
          return <Badge color="danger">greu</Badge>;
      case "medium":
          return <Badge color="warning">mediu</Badge>;
      case "easy":
          return <Badge color="info">usor</Badge>;
      default:
          return <Badge color="success">{ difficulty }</Badge>;
  }
}

export default DifficultyBadge;