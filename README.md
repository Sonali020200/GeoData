# GeoData

# Geospatial Data Management Web Application

This is a fullstack web application built to allow users to manage and visualize geospatial data. The application consists of a backend written in Go and a frontend built with React/Next.js. Users can create an account, upload GeoJSON or KML files, render these files on a map using React Leaflet, draw custom shapes on the map, save those shapes, and edit them.

## Table of Contents

- [Features](#features)
- [Technical Stack](#technical-stack)
- [Getting Started](#getting-started)


## Features

1. User authentication and account management
2. User-friendly interface for uploading and managing GeoJSON and KML files
3. Integration of React Leaflet library for rendering maps and uploaded GeoJSON/KML files
4. Map view where users can draw custom shapes using Leaflet
5. Ability to save drawn shapes to user accounts
6. Functionality to edit existing shapes, including modifying their geometry and attributes

## Technical Stack

### Frontend

- React/Next.js
- React Leaflet
- Leaflet
- Axios for HTTP requests
- CSS for styling

### Backend

- Go programming language
- RESTful APIs for user authentication, account management, and geospatial data management
- MongoDB or PostgreSQL for data storage
- Gorilla Mux for routing

## Getting Started

To run this project locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone <repository-url>
