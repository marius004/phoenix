import Badge from 'components/Badge/Badge.js';

const GradeBadge = ({ grade }) => {
    switch(grade) {
        case "9":
            return <Badge color="rose">clasa {grade}</Badge>;
        case "10":
            return <Badge color="danger">clasa {grade}</Badge>;
        case "11":
            return <Badge color="warning">clasa {grade}</Badge>;
        case "12":
            return <Badge color="info">clasa {grade}</Badge>;
        default:
            return <Badge color="success">{grade}</Badge>;
    }
};



export default GradeBadge;