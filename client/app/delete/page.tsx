'use client'
import { useState } from 'react'
import { Button } from '@mui/material'

const deleteOne = (
	key: string,
	value: string,

) => {
	if (key === '' || value === '' ){
		window.alert('删除参数都不能为空!')
		return
	}

	window.fetch('http://localhost:4000/user/delete', {
		method: 'DELETE',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({
			key,
			value,
		}),
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

export default function DeleteOne () {
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

			<Button onClick={ () => deleteOne(key, value) }>delete</Button>
		</>
	)
};
