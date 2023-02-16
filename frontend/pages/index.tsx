import Head from 'next/head'
import Image from 'next/image'
import { Inter } from '@next/font/google'
const inter = Inter({ subsets: ['latin'] })
import { FormEventHandler, useState } from 'react';

export default function signUp() {
  const [userInfo, setUserInfo] = useState({
    email: '',
    password: ''
  });

  const handleSubmit = event => {
    event.preventDefault();
    
  }
  
  return (
    <>
      <div className=' w-full h-screen flex justify-center items-center'>
        <div className='bg-red-300 w-[600px] h-80 px-24 flex flex-col justify-start items-center'>
          <div className='w-full h-auto flex flex-row justify-center items-center'>
            <Image
                    src="/next.svg"
                    alt="Next.js Logo"
                    width={180}
                    height={37}
                    priority
                  />
          </div>
          <div className='bg-green-300 pt-6'>
            <h2 className='justify-center'>Log in to your notify account</h2>
            <form className='pt-6 space-y-3' onSubmit={handleSubmit}>
                <input
                    placeholder="Enter Email"
                    type="text" className="w-full py-2 px-2" value={userInfo.email} onChange={({ target }) =>
                    setUserInfo({ ...userInfo, email: target.value })
                  }></input>
                  <input
                    placeholder="Password"
                    type="text" className="w-full py-2 px-2" value={userInfo.password} onChange={({ target }) =>
                    setUserInfo({ ...userInfo, password: target.value })
                  }></input>
                  <br></br>
                  <input type="checkbox" name="remember" value="RememberMe"/>
                  <label htmlFor="remember me" className="w-full"> Remember Me</label>
                  <button className="mt-6 mb-6 flex h-14 w-[360px]  flex-row items-center justify-center rounded-lg border border-Neutral-50 bg-white p-4 text-base text-Primary-900"
          >Submit</button>
                  <a href="url" >Forgot Password?</a>
            </form>
          </div> 
        </div>
      </div>
    </>
  )
}
