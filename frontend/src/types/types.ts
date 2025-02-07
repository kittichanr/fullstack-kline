export type AccountFlag = {
  flag_id: number
  flag_type: string
  flag_value: string
}

export type AccountDetail = {
  color: string
  is_main_account: boolean
  progress: number
}

export type Account = {
  account: {
    account_number: string
    currency: string
    dummy_col_3: string
    issuer: string
    type: string
  }
  balance: {
    amount: number
  }
  detail: AccountDetail
  flags: AccountFlag[]
}

export type User = {
  dummy_col_2: string
  greeting: string
  user_id: string
}

export type Banner = {
  banner_id: string
  description: string
  dummy_col_11: string
  image: string
  title: string
  user_id: string
}

export type DebitCard = {
  cards: {
    border_color: string
    color: string
    issuer: string
    name: string
    number: string
    status: string
  }[]
}

export type DataState = {
  user: User | null
  account: Account[] | null
  banner: Banner[] | null
  debitCard: DebitCard[] | null
}
