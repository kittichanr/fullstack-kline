import { ChangeEvent, useEffect, useState } from "react"
import "../index.css"
import { useMutation } from "react-query"
import { useNavigate } from "react-router"
import { authPin } from "../api/auth"

const MAX_LENGTH = 6

const Pin = () => {
  const navigate = useNavigate()

  const [pin, setPin] = useState("")
  const [name, setName] = useState("")

  const [errorMessage, setErrorMessage] = useState("")

  const authMutation = useMutation(
    async ({ pin, name }: { pin: string; name: string }) => {
      const response = await authPin({ pin, name })
      return response.data
    },
    {
      onSuccess: (data) => {
        sessionStorage.setItem("token", data.token)
        alert("Success! Token stored.")
        navigate("/main-bank")
      },
      onError: (error) => {
        console.error(error)
        setErrorMessage("Invalid PIN or User name. Try again.")
      },
    }
  )

  const handleKeyPress = (value: string) => {
    if (pin.length < MAX_LENGTH) {
      setPin(pin + value)
    }
  }

  const handleDelete = () => {
    setPin(pin.slice(0, -1))
    setErrorMessage("")
  }

  const handleInput = (e: ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value)
    setErrorMessage("")
  }

  useEffect(() => {
    if (pin.length === MAX_LENGTH && name.length > 0) {
      authMutation.mutate({ pin, name })
    }
  }, [pin, name])

  return (
    <div className="wrap">
      <main className="container container--pin-type">
        <div className="pin">
          <div className="pin__top">
            <span className="pin__photo">
              <img src="https://dummyimage.com/200x200/999/fff" alt="User" />
            </span>
            <h1 className="pin__name">Interview User</h1>
            <p
              className="pin__dsc"
              style={{
                display: errorMessage ? "block" : "none",
              }}
            >
              {errorMessage}
            </p>
            <div className="pin__dots">
              {[...Array(MAX_LENGTH)].map((_, index) => (
                <span
                  key={index}
                  className={`pin__dot ${
                    index < pin.length ? "is-filled" : ""
                  }`}
                ></span>
              ))}
            </div>
          </div>
          <div className="pin__input">
            <input
              type="text"
              value={name}
              onChange={handleInput}
              placeholder="Enter you name"
            />
          </div>
          <div className="pin__btm">
            <a href="#" className="pin__login">
              Login with user name
            </a>
            <span className="pin__kb">Powered by TestLab</span>
            <div className="pin__keys">
              {[...Array(9)].map((_, index) => (
                <button
                  key={index + 1}
                  type="button"
                  className="pin__key"
                  onClick={() => handleKeyPress((index + 1).toString())}
                >
                  {index + 1}
                </button>
              ))}
              <span className="pin__key pin__key--space"></span>
              <button
                type="button"
                className="pin__key"
                onClick={() => handleKeyPress("0")}
              >
                0
              </button>
              <button
                type="button"
                className="pin__key pin__key--del"
                onClick={handleDelete}
              >
                <span className="">Delete</span>
              </button>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}

export default Pin
