import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter, Routes, Route} from 'react-router-dom'
import { createTheme, ThemeProvider } from '@mui/material'

const theme = createTheme({
    palette: {
        primary: {
            main: '#FFFFFF'
        },
        secondary: {
            main: '#2b2a2b'
        }
    }
})

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <ThemeProvider theme={theme}>
        <BrowserRouter>
            <Routes>
                <Route path='/' element={<App/>}/>
            </Routes>
        </BrowserRouter>
    </ThemeProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
