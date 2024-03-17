import React from 'react';

const Profile = ({ user }) => {
    return (
        <div>
            <h2>User Profile</h2>
            <div>
                <strong>Username:</strong> {user.username}
            </div>
            <div>
                <strong>Email:</strong> {user.email}
            </div>

        </div>
    );
};

export default Profile;
