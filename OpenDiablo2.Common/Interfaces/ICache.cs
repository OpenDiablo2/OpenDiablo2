/*  OpenDiablo 2 - An open source re-implementation of Diablo 2 in C#
 *  
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>. 
 */

using System;
using System.Runtime.Caching;

namespace OpenDiablo2.Common.Interfaces
{
    /// <summary>
    /// Provides access to the cache system.
    /// </summary>
    public interface ICache
    {
        /// <summary>
        /// Gets an item from the cache. If the item does not exist, the value factory will be executed to
        /// generate the value.
        /// </summary>
        /// <typeparam name="T">The return type of the value</typeparam>
        /// <param name="key">The name of the cache key, in the form of Type::X::Y::Z where 
        /// Type is the base type, and x/y/z are unique identifiers for the item.</param>
        /// <param name="valueFactory">A function that returns the correct value if it does not already exist.</param>
        /// <param name="cacheItemPolicy">Pass in a new policy to control how this item is handled. Typically you can leave this null.</param>
        /// <returns>The item requested</returns>
        T AddOrGetExisting<T>(string key, Func<T> valueFactory, CacheItemPolicy cacheItemPolicy = null);
        bool Exists(string key);
        T GetExisting<T>(string key) where T : class, new();
        void Add<T>(string key, T value, CacheItemPolicy cacheItemPolicy = null);
    }
}
