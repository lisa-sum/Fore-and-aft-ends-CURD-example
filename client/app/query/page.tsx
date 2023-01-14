'use client'
import { useState } from 'react'
import { Button } from '@mui/material'

const findOne = (
	key: string,
	value: string,

) => {
	if (key === '' || value === '' ){
		window.alert('删除参数都不能为空!')
		return
	}

	window.fetch(`http://localhost:4000/user/find?key=${key}&value=${value}`, {
		method: 'GET',
		headers: { 'Content-Type': 'application/json' },
	})
	.then((res) => {
		if (!res.ok){
			const { statusText } = res
			throw new Error(statusText)
		}
		return res.json()
	})
	.then(res => {
		console.log(res)
	})
	.catch(err => {
		console.error(err)
	})
}

export default function FindOne () {
	const [key, setKey] = useState<string>('')
	const [value, setValue] = useState<string>('')
	return (
		<>
			<label htmlFor=''>Key:
				<input type='search' onChange={ (e) => setKey(e.currentTarget.value) } />
			</label>
			<label htmlFor=''>Value:
				<input type='search' onChange={ (e) => setValue(e.currentTarget.value) } />
			</label>

			<Button onClick={ () => findOne(key, value) }>Find</Button>
		</>
	)
};
