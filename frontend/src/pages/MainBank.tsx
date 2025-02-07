import { useEffect } from "react"
import "../index.css"
import { fetchUserData } from "../api/user"
import { useNavigate } from "react-router"
import {
  fetchAccountData,
  fetchBannerData,
  fetchDebitCardData,
  getUserId,
  isTokenExpired,
} from "../api"
import { Account, Banner, DebitCard, User } from "../types/types"
import { useQuery } from "react-query"
import { ClipLoader } from "react-spinners"
import { maskCardNumber, thBaht } from "../utils/currency"

const BankMain = () => {
  const navigate = useNavigate()

  useEffect(() => {
    if (isTokenExpired()) {
      navigate("/pin", { replace: true })
    }
  }, [navigate])

  const userId = getUserId()
  const { data: user, isLoading: isUserLoading } = useQuery<User>(
    ["user", userId],
    () => fetchUserData(userId)
  )
  const { data: account, isLoading: isAccountLoading } = useQuery<Account[]>(
    ["account", userId],
    () => fetchAccountData(userId)
  )
  const { data: banner, isLoading: isBannerLoading } = useQuery<Banner[]>(
    ["banner", userId],
    () => fetchBannerData(userId)
  )
  const { data: debitCard, isLoading: isDebitCardLoading } =
    useQuery<DebitCard>(["debit-card", userId], () =>
      fetchDebitCardData(userId)
    )

  const isLoading =
    isUserLoading || isAccountLoading || isBannerLoading || isDebitCardLoading

  if (isLoading) {
    return (
      <div className="center-content">
        <ClipLoader size={50} color="#00f" loading={isLoading} />
      </div>
    )
  }

  const mainAccount = account?.filter((i) => i.detail.is_main_account)
  const subAccount = account?.filter((i) => !i.detail.is_main_account)

  const mapAccountBgColor = (color: string) => {
    switch (color) {
      case "#15bbc7":
        return "is-bluegreen"
      case "#557bf2":
        return "is-deepblue"
      case "#9366ed":
        return "is-purple"
      case "#f88355":
        return "is-orange"
      default:
        return ""
    }
  }
  return (
    <div className="wrap">
      <header className="header ">
        <a href="#" className="header__lft header__menu">
          <span className="blind">Menu</span>
        </a>
        <button type="button" className="header__rgt header__cxl">
          <span className="blind">Cancel</span>
        </button>
      </header>

      <main className="container container--main">
        <div className="content_wrap">
          <div className="main-top">
            <h1 className="main-top__tit main-loading main-loading--order1">
              {user?.greeting}
            </h1>
          </div>

          <div className="main-acc main-acc--large main-loading main-loading--order3 ">
            <div className="main-acc__top">
              <h2 className="main-acc__name">
                {mainAccount?.[0]?.account?.type}
              </h2>
              <span className="main-acc__amount">
                {thBaht.format(mainAccount?.[0]?.balance?.amount ?? 0)}
              </span>
              <span className="main-acc__detail main-acc__detail--num">
                {mainAccount?.[0]?.account?.type}{" "}
                {mainAccount?.[0]?.account?.account_number}
              </span>
              <span className="main-acc__detail">
                Powered by {mainAccount?.[0]?.account?.issuer}
              </span>
            </div>

            <div className="main-acc__bottom">
              <div className="main-acc__ani_box">
                <div className="main-acc__ani__item">
                  <span className="main-acc__ani_img1"></span>
                  <span className="main-acc__ani_img2"></span>
                </div>
                <div className="main-acc__ani__item2">
                  <span className="main-acc__ani_img3"></span>
                </div>
              </div>
              <div className="main-acc__link__box">
                <div className="main-acc__link__item">
                  <a
                    href="#"
                    className="main-acc__link main-acc__link--withdrawal"
                  >
                    Withdrawal
                  </a>
                  <a href="#" className="main-acc__link main-acc__link--qr">
                    QR scan
                  </a>
                  <a
                    href="#"
                    className="main-acc__link main-acc__link--addmoney"
                  >
                    Add money
                  </a>
                </div>
              </div>
            </div>
            <button type="button" className="main-acc__more">
              <span className="blind">More Action</span>
            </button>
            <div className="tooltip " style={{ display: "none" }}>
              <button type="button" className="tooltip__btn-more">
                Set main account
              </button>

              <button type="button" className="tooltip__btn-more">
                Copy account number
              </button>

              <button type="button" className="tooltip__btn-more">
                Edit Name and Color
              </button>
            </div>
            <div
              className="tooltip tooltip--bubble tooltip--right-under"
              style={{ display: "none" }}
            >
              <span className="tooltip__txt">
                Change your main account for
                <br />
                Using transfer, Wallet more easliy
              </span>
            </div>
          </div>

          <div className="rctly__wrap main-loading main-loading--order5">
            <ul className="rctly__lst">
              {banner?.map((item, index) => {
                return (
                  <li key={index} className="rctly__item">
                    <a href="#" className="rctly__link">
                      <span className="rctly__thumb">
                        <img src={item.image} alt={item.description} />
                      </span>
                      <span className="rctly__name">{item.title}</span>
                    </a>
                  </li>
                )
              })}
            </ul>
          </div>

          <div className="debit-swipe__wrap main-loading main-loading--order6">
            <div className="debit-swipe__inner">
              <div className="debit-swipe__lst" style={{ width: "1595px" }}>
                {debitCard?.cards?.slice(0, 3).map((item, index) => {
                  return (
                    <a
                      key={index}
                      href="#"
                      className="debit-swipe__item"
                      style={{
                        color: item.color,
                        backgroundColor: "#d2cfcb",
                        borderColor: item.border_color,
                      }}
                    >
                      <strong className="debit-swipe__name">{item.name}</strong>
                      <span className="debit-swipe__etc debit-swipe__etc--active">
                        <span className="debit-swipe__etc__num">
                          {maskCardNumber(item.number.toString())}
                        </span>
                      </span>
                      <span
                        className="debit-swipe__issue"
                        style={{ color: item.color }}
                      >
                        Issued by {item.issuer}
                      </span>
                    </a>
                  )
                })}

                <a
                  href="#"
                  className="debit-swipe__item debit-swipe__item--all"
                >
                  See all
                </a>
              </div>
            </div>
          </div>

          {subAccount?.map((item, index) => {
            return (
              <div
                key={index}
                className={`main-acc ${mapAccountBgColor(
                  item.detail.color
                )} is-small`}
              >
                <div className="main-acc__top">
                  <h2 className="main-acc__name">{item.account.type}</h2>
                  <span className="main-acc__amount">
                    {thBaht.format(item.balance.amount)}
                  </span>
                  {item.flags.map((flag, index) => {
                    return (
                      <span
                        key={index}
                        className={`main-acc__flag  ${
                          index === 0 && "main-acc__flag--lock"
                        }`}
                      >
                        {flag.flag_value}
                      </span>
                    )
                  })}
                </div>
                <div className="main-acc__bottom">
                  <span className="main-acc__detail">
                    {item.account.type} {item.account.account_number}
                  </span>
                  <span className="main-acc__detail">
                    Powered by {item.account.issuer}
                  </span>
                </div>
                <button
                  type="button"
                  className="main-acc__more main-acc__more--small"
                >
                  <span className="blind">More Action</span>
                </button>

                <div className="main-acc__circle">
                  <svg
                    className="graph-bar"
                    width="100%"
                    height="100%"
                    viewBox="0 0 42 42"
                  >
                    <circle
                      cx="21"
                      cy="21"
                      r="15.91549430918954"
                      fill="transparent"
                      stroke="rgba(0,0,0,0.07)"
                      strokeWidth="1.5"
                    ></circle>
                    <circle
                      className="gauge"
                      cx="21"
                      cy="21"
                      r="15.91549430918954"
                      fill="transparent"
                      stroke="#fff"
                      strokeWidth="1.5"
                      strokeLinecap="round"
                      strokeDashoffset="25"
                      style={{
                        strokeDasharray: `${item.detail.progress} ${
                          100 - item.detail.progress
                        }`,
                      }}
                    ></circle>
                  </svg>
                  <div className="main-acc__num">
                    <span className="percent">{item.detail.progress}</span>
                    <span className="unit">%</span>
                  </div>
                </div>
              </div>
            )
          })}

          <a href="#" className="main-prod">
            <span className="main-prod__cms-ico">
              <img src="https://dummyimage.com/54x54/999/fff" alt="" />
            </span>
            <strong className="main-prod__tit">Want some money?</strong>
            <p className="main-prod__dsc">You can start apply 'Clare'</p>
          </a>

          <div className="main-tb">
            <a href="#" className="link-to">
              Total Balance
            </a>
          </div>
        </div>
      </main>
    </div>
  )
}

export default BankMain
