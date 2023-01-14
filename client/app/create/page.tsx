'use client'
import { Box, Button, TextField } from '@mui/material'
import { useState } from 'react'

import Radio from '../../components/Radio'

const submit = (
	usr: string,
	nickname: string,
	age: number,
	pwd: string,
	rePwd: string,
	tel: string,
) => {
	fetch('http://localhost:4000/user/create', {
		method: 'PUT',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({
			nickname,
			age,
			tel,
			username: usr,
			password: pwd,
			re_password: rePwd,
		}),
	})
	.then(async res => {
		console.log(await res.json())
	})
}

export default function Create () {
	const [usr, setUsr] = useState<string>('')
	const [pwd, setPwd] = useState<string>('')
	const [rePwd, setRePwd] = useState<string>('')
	const [nickname, setNickname] = useState<string>('')
	const [age, setAge] = useState<number>(0)

	const [tel, setTel] = useState<string>('')

	return (

		<Box sx={ {
			width: '300px',
			mx: 'auto',
			height: '80vh',
			display: 'flex',
			justifyContent: 'space-evenly',
			flexDirection: 'column',
		} }>
			<TextField id='username' label='Username' variant='standard'
								 onChange={ (e) => setUsr(e.currentTarget.value) } />
			<TextField id='username' label='Nickname' variant='standard'
								 onChange={ (e) => setNickname(e.currentTarget.value) } />
			<TextField id='password' label='Password' variant='standard'
								 onChange={ (e) => setPwd(e.currentTarget.value) } />
			<TextField id='password' label='RePassword' variant='standard'
								 onChange={ (e) => setRePwd(e.currentTarget.value) } />
			<label htmlFor='age'>
				<input type='radio' min='1' max='100' id='age' onChange={ (e) => setAge(Number(e.currentTarget.value)) } />
			</label>
			<TextField id='username' label='Tel' variant='standard'
								 onChange={ (e) => setTel(e.currentTarget.value) } />

			<Button onClick={ () => submit(usr, nickname, age, pwd, rePwd, tel) }>submit</Button>
		</Box>

	)
};
