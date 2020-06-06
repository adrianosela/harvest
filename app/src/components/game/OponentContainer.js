import React from "react";
import {Image, Row, Col} from "react-bootstrap";
import '../styles/OponentContainer.sass';
import OponentsCards from './OponentsCards'

class OponentContainer extends React.Component {
    render() {
        return(
            <div className="mr-xxl">
                <Row className="zeroMargin">
                    <Image src="https://media-exp1.licdn.com/dms/image/C5603AQEMzM2EEInVdg/profile-displayphoto-shrink_100_100/0?e=1596672000&v=beta&t=5l2pOwdfgDUK3D9vdfv0bco7QHgILaJ6JN18VA1fJbU" className="oponentContainerImage mr-m" roundedCircle></Image>
                    <span>Felipe Ballesteros</span>
                </Row>
                <Row className="zeroMargin">
                    <span>40 pts</span>
                    <OponentsCards />
                </Row>
            </div>
        )
    }

}

export default (OponentContainer);
