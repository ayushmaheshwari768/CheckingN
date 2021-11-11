import React from 'react';
import './Background.css';
import { useParams, Link, useHistory } from 'react-router-dom'

const AppointmentConfirmation = () => {
    const { course } = useParams();
    const { date } = useParams();
    const { time } = useParams();
    const { tutor } = useParams();
    const { location } = useParams();

    const handleConfirm = async (e) => {
        console.log('confirm')
        // e.preventDefault();

        // var request = {
        //     method: 'POST',
        //     headers: {'Content-Type': 'application/json'},
        //     body: JSON.stringify({
        //         student: ,
        //         tutor: ,
        //         course: ,
        //         date: ,
        //         time: ,
        //         location: ,
        //     })
        // }
    }

    // uri: /tutoring/departments/:course/:date/:time/:tutor/:location/confirmation
    // eg. http://localhost:3000/tutoring/departments/CS174/12-23-2021/12:22:00/Nhien%20Lam/online/confirmation
    return (
        <div className='confirmation-container'>
            <h4 className="title">New Appointment</h4>
            <div>
                <p><b>Tutor:</b> {tutor}</p>
                <p><b>Course:</b> {course}</p>
                <p><b>Date:</b> {date}</p>
                <p><b>Time:</b> {time}</p>
                <p><b>Location:</b> {location}</p>
            </div>
            <div className='btn-container'>
                <button className='btn-primary' onClick={handleConfirm}>Confirm</button>
                <Link to={"/tutoring/departments/" + course}>
                    <button className='btn-primary'>Cancel</button>
                </Link>
            </div>
        </div>
    );
};

export default AppointmentConfirmation;