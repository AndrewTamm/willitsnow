import React from 'react';
import './Footer.scss';

class Footer extends React.Component {
  render() {
    return (
      <div className="Footer">
        <p>Created by <a href='https://github.com/AndrewTamm/willitsnow'>Andrew Tamm</a></p>
        <p>Inspired by <a href='https://www.willitsnowinpdx.com/'>Will It Snow In PDX?</a></p>
      </div>
    );
  }
}

export default Footer;
