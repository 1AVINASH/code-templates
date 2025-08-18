import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import AppRoutes from './routes/AppRoutes'
import './App.css'
import {Header} from './features/menu'

function App() {

  return (
    <>
    <div className="bg-red-100">Ok</div>
    <AppRoutes />
    <Header />
    </>
  )
}

export default App
