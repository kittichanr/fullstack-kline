import { useEffect } from "react"
import { useNavigate } from "react-router"
import "../index.css"

export default function Splash() {
  const navigate = useNavigate()

  useEffect(() => {
    const timer = setTimeout(() => {
      navigate("/pin")
    }, 3000)

    return () => clearTimeout(timer)
  }, [navigate])

  return (
    <div className="wrap">
      <div className="splash">
        <div className="loader"></div>
      </div>
    </div>
  )
}
