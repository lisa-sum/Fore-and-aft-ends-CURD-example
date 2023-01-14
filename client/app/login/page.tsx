'use client'
import { useState } from 'react'
import { Button, Box, TextField } from '@mui/material'

const submit = (usr: string, pwd: string) => {
	fetch('http://localhost:4000/login', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({
			username: usr,
			password: pwd,
		}),
	})
	.then(async res => {
		console.log(await res.json())
	})
}

export default function Login () {
	const [usr, setUsr] = useState('')
	const [pwd, setPwd] = useState('')

	return (
		<>

			<Box sx={ {
				width:'60vw',
				mx:'auto',
				display: 'flex',
				justifyContent: 'center',
				flexDirection: 'column',
			} }>

				<TextField id='username' label='Standard' variant='standard'
									 onChange={ (e) => setUsr(e.currentTarget.value) } />
				<TextField id='password' label='Standard' variant='standard'
									 onChange={ (e) => setPwd(e.currentTarget.value) } />
				<Button onClick={ () => submit(usr, pwd) }>submit</Button>
			</Box>
		</>
	)
};
